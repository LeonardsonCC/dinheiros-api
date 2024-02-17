package domain

import "time"

type Account struct {
	ID        int       `db:"account_id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	Color     string    `db:"color" json:"color" binding:"omitempty,hexcolor"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
