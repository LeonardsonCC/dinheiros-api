package categories_repo

import "github.com/LeonardsonCC/dinheiros/categories"

func (c CategoryRepository) ListByUser(userID int) ([]categories.Category, error) {
	var cats []categories.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE user_id = $1 ORDER BY category_id", userID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}
