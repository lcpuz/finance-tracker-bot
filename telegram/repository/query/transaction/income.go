package transactionQuery

const (
	AddUncategorizedIncome = `INSERT INTO tbl_transactions (user_id, transaction_type, transaction_amount, transaction_description) VALUES ($1, 2, $3, $4)`
)
