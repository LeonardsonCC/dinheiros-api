package accounts_repo

func (r AccountRepository) Delete(userID, accountID int) error {
	_, err := r.DB.Exec("DELETE FROM accounts WHERE user_id = $1 AND account_id = $2", userID, accountID)
	if err != nil {
		return err
	}

	return nil
}
