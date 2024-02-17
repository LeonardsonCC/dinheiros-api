package repository

import (
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) Get(userID int) ([]domain.Account, error) {
	var a []domain.Account

	err := r.DB.Select(&a, "SELECT * FROM accounts WHERE user_id = $1 ORDER BY account_id", userID)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (r AccountRepository) Delete(userID, accountID int) error {
	_, err := r.DB.Exec("DELETE FROM accounts WHERE user_id = $1 AND account_id = $2", userID, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (r AccountRepository) Create(u domain.Account) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("INSERT INTO accounts (user_id, name, color) VALUES (:user_id, :name, :color)", u)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r AccountRepository) Update(u domain.Account) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	// where by account id and user id
	// so it won't update accounts from other users
	_, err = tx.NamedExec("UPDATE accounts SET name=:name, color=:color WHERE account_id = :account_id AND user_id = :user_id", u)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
