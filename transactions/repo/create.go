package transactions_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (r TransactionsRepository) Create(t domain.Transaction) (int, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, err
	}

	transactionID := 0
	err = tx.QueryRow(
		"INSERT INTO transactions (account_id, description, value, date, type) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id",
		t.AccountID,
		t.Description,
		t.Value,
		t.Date,
		t.Type,
	).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
