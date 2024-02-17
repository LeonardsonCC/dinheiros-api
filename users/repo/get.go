package users_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (r UserRepository) Get(email string) (domain.User, error) {
	var u domain.User

	err := r.DB.Get(&u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
