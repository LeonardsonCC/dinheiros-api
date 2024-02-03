package accounts

import "time"

type Account struct {
	ID        int       `db:"account_id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
