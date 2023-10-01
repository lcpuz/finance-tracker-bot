package statesRepository

import (
	"github.com/jmoiron/sqlx"
	stateQuery "github.com/lcpuz/finance-tracker-bot/telegram/repository/query/state"
)

type IncomeCategoryStatePostgres struct {
	db *sqlx.DB
}

func NewIncomeCategoryStatePostgres(db *sqlx.DB) *IncomeCategoryStatePostgres {
	return &IncomeCategoryStatePostgres{db: db}
}

func (r *IncomeCategoryStatePostgres) GetIncomeCategoryState(UserID int64) (int8, error) {
	var State int8

	err := r.db.QueryRow(stateQuery.GetIncomeCategoryState, UserID).Scan(&State)
	if err != nil {
		return 0, err
	}

	return State, nil
}

func (r *IncomeCategoryStatePostgres) UpdateIncomeCategoryState(UserID int64, State int8) error {
	_, err := r.db.Exec(stateQuery.UpdateIncomeCategoryState, State, UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *IncomeCategoryStatePostgres) DeleteIncomeCategoryState(UserID int64) error {
	_, err := r.db.Exec(stateQuery.DeleteIncomeCategoryState, UserID)
	if err != nil {
		return err
	}

	return nil
}
