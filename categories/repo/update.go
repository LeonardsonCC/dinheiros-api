package categories_repo

import (
	"github.com/LeonardsonCC/dinheiros/internal/domain"
)

func (c CategoryRepository) Update(cat domain.Category) error {
	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("UPDATE categories SET name = :name WHERE category_id = :category_id", cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
