package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lcpuz/finance-tracker-bot/telegram/repository/query"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(UserRequest request.CreateUsersRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query.CreateUser,
		UserRequest.UserChatID,
		UserRequest.UserName,
		UserRequest.UserNickName,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetUserID(UserChatID int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var UserID int64

	err := r.db.QueryRowContext(ctx, query.GetUserID, UserChatID).Scan(&UserID)
	if err != nil {
		return 0, err
	}

	return UserID, nil
}
