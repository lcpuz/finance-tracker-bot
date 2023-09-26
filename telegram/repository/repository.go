package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type CategoryRepository interface {
	GetCategories(UserID int64) ([]request.CategoriesResponse, error)
	CreateCategory(Category request.CreateCategoryRequest) error
}

type UserRepository interface {
	CreateUser(User request.CreateUsersRequest) error
	GetUserID(UserChatID int64) (int64, error)
}

type Repository struct {
	CategoryRepository
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CategoryRepository: NewCategoryPostgres(db),
		UserRepository:     NewUserPostgres(db),
	}
}
