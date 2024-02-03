package accounts_repo

import "github.com/LeonardsonCC/dinheiros/accounts"

func (r AccountRepository) Create(u accounts.Account) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedQuery("INSERT INTO accounts (user_id, name) VALUES (:user_id, :name)", u)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
