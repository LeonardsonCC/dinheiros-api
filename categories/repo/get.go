package categories_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (c CategoryRepository) Get(categoryID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE category_id = $1 ORDER BY category_id", categoryID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}
