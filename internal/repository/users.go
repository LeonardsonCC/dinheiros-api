package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r UserRepository) List() ([]domain.User, error) {
	var u []domain.User

	err := r.DB.Select(&u, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r UserRepository) Get(ctx context.Context, email string) (domain.User, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, "repository user")
	defer sp.End()

	var u domain.User

	err := r.DB.GetContext(ctx, &u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r UserRepository) Create(ctx context.Context, u domain.User) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, "repository user")
	defer sp.End()

	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.QueryxContext(ctx, "INSERT INTO users (email) VALUES ($1)", u.Email)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
