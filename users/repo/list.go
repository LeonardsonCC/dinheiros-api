package users_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (r UserRepository) List() ([]domain.User, error) {
	var u []domain.User

	err := r.DB.Select(&u, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return u, nil
}
