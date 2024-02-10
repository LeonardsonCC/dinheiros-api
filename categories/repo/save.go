package categories_repo

import (
	"github.com/LeonardsonCC/dinheiros/categories"
)

func (r CategoriesRepository) Save(userID, transactionID int, cats []string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	err = r.DeleteAllCategories(transactionID)
	if err != nil {
		return err
	}

	catsToInsert := make([]categories.Category, 0, len(cats))
	for _, c := range cats {
		catsToInsert = append(catsToInsert, categories.Category{
			Name:          c,
			TransactionID: transactionID,
			UserID:        userID,
		})
	}

	_, err = tx.NamedExec("INSERT INTO transaction_category (category_name, transaction_id, user_id) VALUES (:category_name, :transaction_id, :user_id)", catsToInsert)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (r CategoriesRepository) DeleteAllCategories(transactionID int) error {
	_, err := r.DB.Exec("DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)
	if err != nil {
		return err
	}

	return nil
}
