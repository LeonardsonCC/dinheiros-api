package accounts_repo

import "github.com/LeonardsonCC/dinheiros/internal/domain"

func (r AccountRepository) Get(userID int) ([]domain.Account, error) {
	var a []domain.Account

	err := r.DB.Select(&a, "SELECT * FROM accounts WHERE user_id = $1 ORDER BY account_id", userID)
	if err != nil {
		return a, err
	}

	return a, nil
}
