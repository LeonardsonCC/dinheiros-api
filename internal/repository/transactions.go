package repository

import (
	"context"

	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry/spans"
	"github.com/jmoiron/sqlx"
)

type TransactionsRepository struct {
	DB *sqlx.DB
}

func (r TransactionsRepository) List(ctx context.Context, userID, accountID int) ([]domain.Transaction, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.TransactionRepository)
	defer sp.End()

	var t []domain.Transaction

	query := "SELECT t.* FROM transactions t JOIN accounts a ON (a.account_id = t.account_id) WHERE a.user_id = $1"
	if accountID != 0 {
		query += " AND t.account_id = $2"
	}
	query += " ORDER BY t.transaction_id"

	params := make([]any, 0, 2)
	params = append(params, userID)
	if accountID != 0 {
		params = append(params, accountID)
	}

	err := r.DB.SelectContext(ctx, &t, query, params...)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r TransactionsRepository) Get(ctx context.Context, transactionID int) ([]domain.Transaction, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.TransactionRepository)
	defer sp.End()

	var t []domain.Transaction

	err := r.DB.SelectContext(ctx, &t, "SELECT * FROM transactions WHERE transaction_id = $1 ORDER BY transaction_id", transactionID)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (r TransactionsRepository) Delete(ctx context.Context, transactionID int) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.TransactionRepository)
	defer sp.End()

	_, err := r.DB.ExecContext(ctx, "DELETE FROM transactions WHERE transaction_id = $1", transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (r TransactionsRepository) Create(ctx context.Context, t domain.Transaction) (int, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.TransactionRepository)
	defer sp.End()

	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, err
	}

	transactionID := 0
	err = tx.QueryRowContext(
		ctx,
		"INSERT INTO transactions (account_id, description, value, date, type) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id",
		t.AccountID,
		t.Description,
		t.Value,
		t.Date,
		t.Type,
	).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}

func (r TransactionsRepository) Update(ctx context.Context, t domain.Transaction) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.TransactionRepository)
	defer sp.End()

	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	query, err := tx.PrepareNamedContext(ctx, "UPDATE transactions SET description=:description, value=:value, date=:date, type=:type, account_id=:account_id WHERE transaction_id = :transaction_id")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, t)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
