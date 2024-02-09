package transactions_repo

import "github.com/LeonardsonCC/dinheiros/transactions"

func (r TransactionsRepository) Update(t transactions.Transaction) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("UPDATE transactions SET description=:description, value=:value, date=:date, type=:type, account_id=:account_id WHERE transaction_id = :transaction_id", t)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
