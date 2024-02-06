package transactions_repo

import "github.com/jmoiron/sqlx"

type TransactionsRepository struct {
	DB *sqlx.DB
}
