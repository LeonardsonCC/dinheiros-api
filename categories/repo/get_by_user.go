package categories_repo

import "github.com/LeonardsonCC/dinheiros/categories"

func (r CategoriesRepository) GetUserTransactionsCategories(userID int) (map[int][]categories.Category, error) {
	var cats []categories.Category

	err := r.DB.Select(&cats, "SELECT * FROM transaction_category WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	catsByUser := make(map[int][]categories.Category)
	for _, c := range cats {
		catsByUser[c.TransactionID] = append(catsByUser[c.TransactionID], c)
	}

	return catsByUser, nil
}
