package transactions_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (r TransactionsRepository) Get(transactionID int) ([]domain.Transaction, error) {
	var t []domain.Transaction

	err := r.DB.Select(&t, "SELECT * FROM transactions WHERE transaction_id = $1 ORDER BY transaction_id", transactionID)
	if err != nil {
		return t, err
	}

	return t, nil
}
