package domain

import (
	"time"
)

type TransactionType int8

const (
	INCOME TransactionType = iota
	OUTCOME
)

type Transaction struct {
	ID          int             `db:"transaction_id"`
	UserID      int             `db:"-"`
	AccountID   int             `db:"account_id"`
	Description string          `db:"description" binding:"required"`
	Value       int             `db:"value" binding:"required"`
	Date        time.Time       `db:"date" binding:"required"`
	Type        TransactionType `db:"type" binding:"required,oneof=income outcome"`
	CreatedAt   time.Time       `db:"created_at"`
}

type TransactionJson struct {
	ID          int            `json:"id"`
	AccountID   int            `json:"account_id"`
	Description string         `json:"description" binding:"required,max=300"`
	Value       float64        `json:"value" binding:"required"`
	Date        string         `json:"date" binding:"required"`
	Type        string         `json:"type" binding:"required,oneof=INCOME OUTCOME"`
	Categories  []CategoryJson `json:"categories"`
	CreatedAt   time.Time      `json:"created_at"`
}

func MapJsonToDomain(in TransactionJson) (Transaction, []Category, error) {
	value := int(in.Value * 100)
	date, err := time.Parse("02/01/2006", in.Date)
	if err != nil {
		return Transaction{}, nil, err
	}

	t := INCOME
	switch in.Type {
	case "INCOME":
		t = INCOME
	case "OUTCOME":
		t = OUTCOME
	}

	cats := make([]Category, 0, len(in.Categories))
	for _, v := range in.Categories {
		cats = append(cats, Category{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return Transaction{
		ID:          in.ID,
		AccountID:   in.AccountID,
		Description: in.Description,
		Value:       value,
		Date:        date,
		Type:        t,
	}, cats, nil
}

func MapDomainToJson(in Transaction, cats []Category) TransactionJson {
	value := float64(in.Value) / float64(100)
	date := in.Date.Format("2006-01-02")

	var t string
	switch in.Type {
	case INCOME:
		t = "INCOME"
	case OUTCOME:
		t = "OUTCOME"
	}

	cs := make([]CategoryJson, 0, len(cats))
	for _, c := range cats {
		cs = append(cs, CategoryJson{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return TransactionJson{
		ID:          in.ID,
		AccountID:   in.AccountID,
		Description: in.Description,
		Value:       value,
		Date:        date,
		Type:        t,
		Categories:  cs,
		CreatedAt:   in.CreatedAt,
	}
}
