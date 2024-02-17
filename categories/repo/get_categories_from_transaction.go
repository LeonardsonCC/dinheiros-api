package categories_repo

import (
	"github.com/LeonardsonCC/dinheiros/internal/domain"
)

func (c CategoryRepository) GetCategoriesFromTransaction(transactionID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.Select(&cats, "SELECT c.* FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) WHERE tc.transaction_id = $1 ORDER BY category_id", transactionID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

func (c CategoryRepository) GetCategoriesFromAccount(userID, accountID int) (map[int][]domain.Category, error) {
	var cats []domain.Category

	var param int
	var query string
	if userID > 0 {
		query = "SELECT c.*, tc.transaction_id FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) WHERE c.user_id = $1"
		param = userID
	}
	if accountID > 0 {
		query = "SELECT c.*, tc.transaction_id, t.account_id FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) JOIN transactions t ON (t.transaction_id = tc.transaction_id) WHERE t.account_id = $1"
		param = accountID
	}

	err := c.DB.Select(&cats, query, param)
	if err != nil {
		return nil, err
	}

	cs := make(map[int][]domain.Category)
	for _, v := range cats {
		cs[v.TransactionID] = append(cs[v.TransactionID], v)
	}

	return cs, nil
}
