package repository

import (
	"github.com/jmoiron/sqlx"
	incomeCategoryRepository "github.com/lcpuz/finance-tracker-bot/telegram/repository/categories"
	statesRepository "github.com/lcpuz/finance-tracker-bot/telegram/repository/states"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type IncomeCategoryRepository interface {
	GetCategories(UserID int64) ([]request.CategoriesResponse, error)
	CreateCategory(Category request.CreateCategoryRequest) error
}

type StatesRepository interface {
	CreateDefaultStates(UserID int64) error
}

type IncomeCategoryStateRepository interface {
	CreateIncomeCategoryState(UserID int64, State int8) error
	GetIncomeCategoryState(UserID int64) (int8, error)
	UpdateIncomeCategoryState(UserID int64, State int8) error
	DeleteIncomeCategoryState(UserID int64) error
}

type UserRepository interface {
	CreateUser(User request.CreateUsersRequest) error
	GetUserID(UserChatID int64) (int64, error)
}

type Repository struct {
	UserRepository
	StatesRepository
	IncomeCategoryRepository
	IncomeCategoryStateRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:                NewUserPostgres(db),
		StatesRepository:              statesRepository.NewStatesRepository(db),
		IncomeCategoryRepository:      incomeCategoryRepository.NewCategoryPostgres(db),
		IncomeCategoryStateRepository: statesRepository.NewIncomeCategoryStatePostgres(db),
	}
}
