package query

const (
	// CreateCategory ...
	CreateCategory = `INSERT INTO tbl_categories (name, user_id) VALUES ($1, $2)`

	// GetCategories ...
	GetCategories = `SELECT id, name FROM tbl_categories WHERE user_id = $1`
)
