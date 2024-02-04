package accounts_repo

import "github.com/LeonardsonCC/dinheiros/accounts"

func (r AccountRepository) Get(userID int) ([]accounts.Account, error) {
	var a []accounts.Account

	err := r.DB.Select(&a, "SELECT * FROM accounts WHERE user_id = $1", userID)
	if err != nil {
		return a, err
	}

	return a, nil
}
