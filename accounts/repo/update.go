package accounts_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

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
