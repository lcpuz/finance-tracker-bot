package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lcpuz/finance-tracker-bot/telegram/repository/query"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategory(CategoryRequest request.CreateCategoryRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query.CreateCategory, CategoryRequest.CategoryName, CategoryRequest.CategoryUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryPostgres) GetCategories(UserID int64) ([]request.CategoriesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Categories []request.CategoriesResponse

	row, err := r.db.QueryContext(ctx, query.GetCategories, UserID)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var Category request.CategoriesResponse
		err = row.Scan(&Category.CategoryID, &Category.CategoryName)
		if err != nil {
			return nil, err
		}

		Categories = append(Categories, Category)
	}

	return Categories, nil
}
