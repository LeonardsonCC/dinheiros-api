package repository

import (
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func (c CategoryRepository) Get(categoryID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE category_id = $1 ORDER BY category_id", categoryID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

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

func (c CategoryRepository) ListByUser(userID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE user_id = $1 ORDER BY category_id", userID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

func (c CategoryRepository) Delete(categoryID int) error {
	_, err := c.DB.Exec("DELETE FROM categories WHERE category_id = $1", categoryID)

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) DeleteByTransaction(transactionID int) error {
	_, err := c.DB.Exec("DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Create(cat domain.Category) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("INSERT INTO categories (user_id, name) VALUES (:user_id, :name)", cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Update(cat domain.Category) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("UPDATE categories SET name = :name WHERE category_id = :category_id", cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) AddCategoryToTransaction(transactionID int, cats []domain.Category) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	err = c.deleteAllCategoriesFromTransaction(transactionID)
	if err != nil {
		return err
	}

	if len(cats) == 0 {
		return nil
	}

	tcs := make([]domain.TransactionCategory, 0, len(cats))
	for _, cc := range cats {
		tcs = append(tcs, domain.TransactionCategory{
			TransactionID: transactionID,
			CategoryID:    cc.ID,
		})
	}

	_, err = tx.NamedExec("INSERT INTO transaction_category (category_id, transaction_id) VALUES (:category_id, :transaction_id)", tcs)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) deleteAllCategoriesFromTransaction(transactionID int) error {
	_, err := c.DB.Exec("DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)

	return err
}
