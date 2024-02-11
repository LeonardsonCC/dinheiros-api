package categories

type Category struct {
	ID            int    `db:"category_id" json:"id"`
	UserID        int    `db:"user_id" json:"-"`
	Name          string `db:"name" json:"name"`
	TransactionID int    `db:"transaction_id" json:"-"`
	AccountID     int    `db:"account_id" json:"-"`
}

type TransactionCategory struct {
	TransactionID int `db:"transaction_id"`
	CategoryID    int `db:"category_id"`
}

type CategoryJson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
