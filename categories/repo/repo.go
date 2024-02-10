package categories_repo

import "github.com/jmoiron/sqlx"

type CategoriesRepository struct {
	DB *sqlx.DB
}
