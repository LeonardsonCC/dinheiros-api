package categories_repo

import "github.com/jmoiron/sqlx"

type CategoryRepository struct {
	DB *sqlx.DB
}
