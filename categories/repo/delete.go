package categories_repo

func (c CategoryRepository) Delete(categoryID int) error {
	_, err := c.DB.Exec("DELETE FROM categories WHERE category_id = $1", categoryID)

	if err != nil {
		return err
	}

	return nil
}
