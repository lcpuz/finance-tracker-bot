package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type CategoryRepository interface {
	GetCategories() ([]request.CategoriesResponse, error)
	CreateCategory(Category request.CreateCategoryRequest) error
}

type Repository struct {
	CategoryRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CategoryRepository: NewCategoryPostgres(db),
	}
}
