package stateQuery

const (
	GetUncategorizedIncomeState    = `SELECT income_state FROM tbl_user_states WHERE user_id = $1`
	UpdateUncategorizedIncomeState = `UPDATE tbl_user_states SET income_state = $1 WHERE user_id = $2`
	DeleteUncategorizedIncomeState = `UPDATE tbl_user_states SET income_state = 0 WHERE user_id = $1`
)
