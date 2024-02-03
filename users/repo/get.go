package users_repo

import "github.com/LeonardsonCC/dinheiros/users"

func (r UserRepository) Get(email string) (users.User, error) {
	var u users.User

	err := r.DB.Get(&u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return users.User{}, err
	}

	return u, nil
}
