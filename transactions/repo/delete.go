package transactions_repo

func (r TransactionsRepository) Delete(transactionID int) error {
	_, err := r.DB.Exec("DELETE FROM transactions WHERE transaction_id = $1", transactionID)
	if err != nil {
		return err
	}

	return nil
}
