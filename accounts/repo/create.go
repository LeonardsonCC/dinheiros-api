package accounts_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

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
