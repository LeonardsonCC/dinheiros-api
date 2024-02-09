package transactions_repo

import "github.com/LeonardsonCC/dinheiros/transactions"

func (r TransactionsRepository) List(userID, accountID int) ([]transactions.Transaction, error) {
	var t []transactions.Transaction

	query := "SELECT t.* FROM transactions t JOIN accounts a ON (a.account_id = t.account_id) WHERE a.user_id = $1"
	if accountID != 0 {
		query += " AND t.account_id = $2"
	}
	query += " ORDER BY t.transaction_id"

	params := make([]interface{}, 0, 2)
	params = append(params, userID)
	if accountID != 0 {
		params = append(params, accountID)
	}

	err := r.DB.Select(&t, query, params...)
	if err != nil {
		return t, err
	}

	return t, nil
}
