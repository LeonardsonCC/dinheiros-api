package users_repo

import "github.com/LeonardsonCC/dinheiros/users"

func (r UserRepository) Create(u users.User) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedQuery("INSERT INTO users (email) VALUES (:email)", u)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
