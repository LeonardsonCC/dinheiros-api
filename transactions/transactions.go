package transactions

import (
	"math"
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
	ID          int       `json:"id"`
	Description string    `json:"description" binding:"required,max=300"`
	Value       float64   `json:"value" binding:"required"`
	Date        string    `json:"date" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=INCOME OUTCOME"`
	CreatedAt   time.Time `json:"created_at"`
}

func MapJsonToDomain(in TransactionJson) (Transaction, error) {
	value := int(in.Value * 100)
	date, err := time.Parse("02/01/2006", in.Date)
	if err != nil {
		return Transaction{}, err
	}

	t := INCOME
	switch in.Type {
	case "INCOME":
		t = INCOME
	case "OUTCOME":
		t = OUTCOME
	}

	return Transaction{
		ID:          in.ID,
		Description: in.Description,
		Value:       value,
		Date:        date,
		Type:        t,
	}, nil
}

func MapDomainToJson(in Transaction) TransactionJson {
	value := math.Floor(float64(in.Value) / float64(100))
	date := in.Date.Format("2006-01-02")

	var t string
	switch in.Type {
	case INCOME:
		t = "INCOME"
	case OUTCOME:
		t = "OUTCOME"
	}

	return TransactionJson{
		ID:          in.ID,
		Description: in.Description,
		Value:       value,
		Date:        date,
		Type:        t,
		CreatedAt:   in.CreatedAt,
	}
}
