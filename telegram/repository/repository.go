package repository

import (
	"github.com/jmoiron/sqlx"
	incomeCategoryRepository "github.com/lcpuz/finance-tracker-bot/telegram/repository/categories"
	statesRepository "github.com/lcpuz/finance-tracker-bot/telegram/repository/states"
	transactionsRepository "github.com/lcpuz/finance-tracker-bot/telegram/repository/transactions"
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
	GetIncomeCategoryState(UserID int64) (int8, error)
	UpdateIncomeCategoryState(UserID int64, State int8) error
	DeleteIncomeCategoryState(UserID int64) error
}

type UncategorizedIncomeStateRepository interface {
	GetUncategorizedIncomeState(UserID int64) (int8, error)
	UpdateUncategorizedIncomeState(UserID int64, State int8) error
	DeleteUncategorizedIncomeState(UserID int64) error
}

type IncomeRepository interface {
	AddUncategorizedIncome(UncategorizedIncome request.AddUncategorizedIncomeRequest) error
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
	UncategorizedIncomeStateRepository
	IncomeRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:                     NewUserPostgres(db),
		StatesRepository:                   statesRepository.NewStatesRepository(db),
		IncomeCategoryRepository:           incomeCategoryRepository.NewCategoryPostgres(db),
		IncomeCategoryStateRepository:      statesRepository.NewIncomeCategoryStatePostgres(db),
		UncategorizedIncomeStateRepository: statesRepository.NewUncategorizedIncomePostgres(db),
		IncomeRepository:                   transactionsRepository.NewIncomePostgres(db),
	}
}
