package request

type AddUncategorizedIncomeRequest struct {
	UserID                 int64  `json:"user_id"`
	TransactionAmmount     string `json:"transaction_ammount"`
	TransactionDescription string `json:"transaction_description"`
}
