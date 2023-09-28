package request

type AddUncategorizedIncomeRequest struct {
	UserID                 int64   `json:"user_id"`
	TransactionAmount      float64 `json:"transaction_amount"`
	TransactionDescription string  `json:"transaction_description"`
}
