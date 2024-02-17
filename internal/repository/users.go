package repository

import (
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r UserRepository) List() ([]domain.User, error) {
	var u []domain.User

	err := r.DB.Select(&u, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r UserRepository) Get(email string) (domain.User, error) {
	var u domain.User

	err := r.DB.Get(&u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r UserRepository) Create(u domain.User) error {
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
