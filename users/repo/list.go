package users_repo

import "github.com/LeonardsonCC/dinheiros/users"

func (r UserRepository) List() ([]users.User, error) {
	var u []users.User

	err := r.DB.Select(&u, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return u, nil
}
