package domain

import "time"

type User struct {
	ID        int       `db:"user_id" json:"id"`
	Email     string    `db:"email" json:"email" binding:"required,email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
