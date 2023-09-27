package query

const (
	// CreateCategory ...
	CreateCategory = `INSERT INTO tbl_categories (category_name, user_id) VALUES ($1, $2)`

	// GetCategories ...
	GetCategories = `SELECT category_id, category_name FROM tbl_categories WHERE user_id = $1`
)
