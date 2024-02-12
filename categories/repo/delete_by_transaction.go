package categories_repo

func (c CategoryRepository) DeleteByTransaction(transactionID int) error {
	_, err := c.DB.Exec("DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)

	if err != nil {
		return err
	}

	return nil
}
