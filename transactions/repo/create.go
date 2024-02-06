package transactions_repo

import (
	"github.com/LeonardsonCC/dinheiros/transactions"
)

func (r TransactionsRepository) Create(t transactions.Transaction) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("INSERT INTO transactions (account_id, description, value, date, type) VALUES (:account_id, :description, :value, :date, :type)", t)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
