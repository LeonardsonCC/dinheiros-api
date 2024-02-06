package transactions_repo

import "github.com/LeonardsonCC/dinheiros/transactions"

func (r TransactionsRepository) List(accountID int) ([]transactions.Transaction, error) {
	var t []transactions.Transaction

	err := r.DB.Select(&t, "SELECT * FROM transactions WHERE account_id = $1 ORDER BY transaction_id", accountID)
	if err != nil {
		return t, err
	}

	return t, nil
}
