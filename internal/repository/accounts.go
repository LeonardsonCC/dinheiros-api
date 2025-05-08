package repository

import (
	"context"

	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry/spans"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) Get(ctx context.Context, userID int) ([]domain.Account, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.AccountRepository)
	defer sp.End()

	var a []domain.Account

	err := r.DB.SelectContext(ctx, &a, "SELECT * FROM accounts WHERE user_id = $1 ORDER BY account_id", userID)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (r AccountRepository) Delete(ctx context.Context, userID, accountID int) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.AccountRepository)
	defer sp.End()

	_, err := r.DB.ExecContext(ctx, "DELETE FROM accounts WHERE user_id = $1 AND account_id = $2", userID, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (r AccountRepository) Create(ctx context.Context, u domain.Account) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.AccountRepository)
	defer sp.End()

	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	query, err := tx.PrepareNamedContext(ctx, "INSERT INTO accounts (user_id, name, color) VALUES (:user_id, :name, :color)")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, u)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r AccountRepository) Update(ctx context.Context, a domain.Account) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	// where by account id and user id
	// so it won't update accounts from other users
	query, err := tx.PrepareNamedContext(ctx, "UPDATE accounts SET name=:name, color=:color WHERE account_id = :account_id AND user_id = :user_id")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, a)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
