package categories

type Category struct {
	Name          string `db:"category_name"`
	TransactionID int    `db:"transaction_id"`
	UserID        int    `db:"user_id"`
}
