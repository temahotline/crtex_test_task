package schemas

type CreateTransaction struct {
	UserID          int64  `json:"user_id" binding:"required"`
	Amount          int64  `json:"amount" binding:"required"`
	BalanceBefore   int64  `json:"balance_before" binding:"required"`
	TransactionType string `json:"transaction_type" binding:"required"`
}

type Transaction struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	Amount          int64  `json:"amount"`
	BalanceBefore   int64  `json:"balance_before"`
	TransactionType string `json:"transaction_type"`
	CreatedAt       string `json:"created_at"`
}
