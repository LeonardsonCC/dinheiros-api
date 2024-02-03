package accounts_repo

import "github.com/jmoiron/sqlx"

type AccountRepository struct {
	DB *sqlx.DB
}
