package transactions_repo

import "github.com/LeonardsonCC/dinheiros/transactions"

func (r TransactionsRepository) Get(transactionID int) ([]transactions.Transaction, error) {
	var t []transactions.Transaction

	err := r.DB.Select(&t, "SELECT * FROM transactions WHERE transaction_id = $1 ORDER BY transaction_id", transactionID)
	if err != nil {
		return t, err
	}

	return t, nil
}
