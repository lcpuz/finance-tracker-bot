package statesRepository

import (
	"github.com/jmoiron/sqlx"
	stateQuery "github.com/lcpuz/finance-tracker-bot/telegram/repository/query/state"
)

type UncategorizedIncomePostgres struct {
	db *sqlx.DB
}

func NewUncategorizedIncomePostgres(db *sqlx.DB) *UncategorizedIncomePostgres {
	return &UncategorizedIncomePostgres{db: db}
}

func (r *UncategorizedIncomePostgres) GetUncategorizedIncomeState(UserID int64) (int8, error) {
	var State int8

	err := r.db.QueryRow(stateQuery.GetUncategorizedIncomeState, UserID).Scan(&State)
	if err != nil {
		return 0, err
	}

	return State, nil
}

func (r *UncategorizedIncomePostgres) UpdateUncategorizedIncomeState(UserID int64, State int8) error {
	_, err := r.db.Exec(stateQuery.UpdateUncategorizedIncomeState, State, UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UncategorizedIncomePostgres) DeleteUncategorizedIncomeState(UserID int64) error {
	_, err := r.db.Exec(stateQuery.DeleteUncategorizedIncomeState, UserID)
	if err != nil {
		return err
	}

	return nil
}
