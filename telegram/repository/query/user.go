package query

const (
	CreateUser = `INSERT INTO tbl_users (user_chat_id, user_name, user_nickname) VALUES ($1, $2, $3)`
	GetUserID  = `SELECT user_id FROM tbl_users WHERE user_chat_id = $1`
)
