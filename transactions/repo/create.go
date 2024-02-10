package transactions_repo

import (
	"github.com/LeonardsonCC/dinheiros/transactions"
)

func (r TransactionsRepository) Create(t transactions.Transaction) (int, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, err
	}

	lastInsertID := 0
	err = tx.QueryRow("INSERT INTO transactions (account_id, description, value, date, type) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id", t.AccountID, t.Description, t.Value, t.Date, t.Type).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
