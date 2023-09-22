package request

type CreateCategoryRequest struct {
	CategoryName   string `json:"category_name"`
	CategoryUserID int    `json:"user_id"`
}

type CategoriesResponse struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
