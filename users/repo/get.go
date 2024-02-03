package users_repo

import "github.com/LeonardsonCC/dinheiros/users"

func (r UserRepository) Get(userID int) (users.User, error) {
	var u users.User

	err := r.DB.Get(&u, "SELECT * FROM users WHERE user_id = $1", userID)
	if err != nil {
		return users.User{}, err
	}

	return u, nil
}
