package transactions_repo

import "github.com/LeonardsonCC/dinheiros/transactions"

func (r TransactionsRepository) Get(accountID int, transactionID int) ([]transactions.Transaction, error) {
	var t []transactions.Transaction

	err := r.DB.Select(&t, "SELECT * FROM transactions WHERE account_id = $1 AND transaction_id = $2 ORDER BY transaction_id", accountID, transactionID)
	if err != nil {
		return t, err
	}

	return t, nil
}
