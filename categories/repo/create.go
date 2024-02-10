package categories_repo

import (
	"github.com/LeonardsonCC/dinheiros/categories"
)

func (c CategoryRepository) Create(cat categories.Category) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("INSERT INTO categories (user_id, name) VALUES (:user_id, :name)", cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
