package categories_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (c CategoryRepository) ListByUser(userID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.Select(&cats, "SELECT * FROM categories WHERE user_id = $1 ORDER BY category_id", userID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}
