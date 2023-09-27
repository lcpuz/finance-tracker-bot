package stateQyery

const (
	CreateIncomeCategoryState = `INSERT INTO tbl_user_states (user_id, income_category_state) VALUES ($1, $2)`
	GetIncomeCategoryState    = `SELECT income_category_state FROM tbl_user_states WHERE user_id = $1`
	UpdateIncomeCategoryState = `UPDATE tbl_user_states SET income_category_state = $1 WHERE user_id = $2`
	DeleteIncomeCategoryState = `UPDATE tbl_user_states SET income_category_state = 0 WHERE user_id = $1`
)
