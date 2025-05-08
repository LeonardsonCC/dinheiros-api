package repository

import (
	"context"

	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry/spans"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func (c CategoryRepository) Get(ctx context.Context, categoryID int) ([]domain.Category, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	var cats []domain.Category

	err := c.DB.SelectContext(ctx, &cats, "SELECT * FROM categories WHERE category_id = $1 ORDER BY category_id", categoryID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

func (c CategoryRepository) GetCategoriesFromTransaction(ctx context.Context, transactionID int) ([]domain.Category, error) {
	var cats []domain.Category

	err := c.DB.SelectContext(ctx, &cats, "SELECT c.* FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) WHERE tc.transaction_id = $1 ORDER BY category_id", transactionID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

func (c CategoryRepository) GetCategoriesFromAccount(ctx context.Context, userID, accountID int) (map[int][]domain.Category, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, "repository transaction")
	defer sp.End()

	var cats []domain.Category

	var param int
	var query string
	if userID > 0 {
		query = "SELECT c.*, tc.transaction_id FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) WHERE c.user_id = $1"
		param = userID
	}
	if accountID > 0 {
		query = "SELECT c.*, tc.transaction_id, t.account_id FROM categories c JOIN transaction_category tc ON (c.category_id = tc.category_id) JOIN transactions t ON (t.transaction_id = tc.transaction_id) WHERE t.account_id = $1"
		param = accountID
	}

	err := c.DB.SelectContext(ctx, &cats, query, param)
	if err != nil {
		return nil, err
	}

	cs := make(map[int][]domain.Category)
	for _, v := range cats {
		cs[v.TransactionID] = append(cs[v.TransactionID], v)
	}

	return cs, nil
}

func (c CategoryRepository) ListByUser(ctx context.Context, userID int) ([]domain.Category, error) {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	var cats []domain.Category

	err := c.DB.SelectContext(ctx, &cats, "SELECT * FROM categories WHERE user_id = $1 ORDER BY category_id", userID)
	if err != nil {
		return cats, err
	}

	return cats, nil
}

func (c CategoryRepository) Delete(ctx context.Context, categoryID int) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	_, err := c.DB.ExecContext(ctx, "DELETE FROM categories WHERE category_id = $1", categoryID)
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) DeleteByTransaction(ctx context.Context, transactionID int) error {
	_, err := c.DB.ExecContext(ctx, "DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Create(ctx context.Context, cat domain.Category) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	query, err := tx.PrepareNamedContext(ctx, "INSERT INTO categories (user_id, name) VALUES (:user_id, :name)")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Update(ctx context.Context, cat domain.Category) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	query, err := tx.PrepareNamedContext(ctx, "UPDATE categories SET name = :name WHERE category_id = :category_id")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, cat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) AddCategoryToTransaction(ctx context.Context, transactionID int, cats []domain.Category) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	tx, err := c.DB.Beginx()
	if err != nil {
		return err
	}

	err = c.deleteAllCategoriesFromTransaction(ctx, transactionID)
	if err != nil {
		return err
	}

	if len(cats) == 0 {
		return nil
	}

	tcs := make([]domain.TransactionCategory, 0, len(cats))
	for _, cc := range cats {
		tcs = append(tcs, domain.TransactionCategory{
			TransactionID: transactionID,
			CategoryID:    cc.ID,
		})
	}

	query, err := tx.PrepareNamedContext(ctx, "INSERT INTO transaction_category (category_id, transaction_id) VALUES (:category_id, :transaction_id)")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, tcs)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) deleteAllCategoriesFromTransaction(ctx context.Context, transactionID int) error {
	ctx, sp := telemetry.GetAppTracer().Start(ctx, spans.CategoryRepository)
	defer sp.End()

	_, err := c.DB.ExecContext(ctx, "DELETE FROM transaction_category WHERE transaction_id = $1", transactionID)

	return err
}
