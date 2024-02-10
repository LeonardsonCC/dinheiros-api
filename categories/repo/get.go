package categories_repo

import "github.com/LeonardsonCC/dinheiros/categories"

func (r CategoriesRepository) GetTransactionCategories(transactionID int) ([]categories.Category, error) {
	var categories []categories.Category

	err := r.DB.Select(&categories, "SELECT * FROM transaction_category WHERE transaction_id = $1", transactionID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
