package categories_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

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
