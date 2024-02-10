package categories_repo

import "github.com/LeonardsonCC/dinheiros/categories"

func (c CategoryRepository) Get(categoryID int) ([]categories.Category, error) {
	var cats []categories.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE category_id = $1 ORDER BY category_id", categoryID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}
