package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry/spans"
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
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.UserRepository)
	defer sp.End()

	var u domain.User

	err := r.DB.GetContext(ctx, &u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return u, nil
}

func (r UserRepository) Create(ctx context.Context, u domain.User) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.UserRepository)
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
