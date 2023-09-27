package statesRepository

import (
	"github.com/jmoiron/sqlx"
	stateQyery "github.com/lcpuz/finance-tracker-bot/telegram/repository/query/state"
)

type StatesRepository struct {
	db *sqlx.DB
}

func NewStatesRepository(db *sqlx.DB) *StatesRepository {
	return &StatesRepository{db: db}
}

func (r *StatesRepository) CreateDefaultStates(UserID int64) error {
	_, err := r.db.Exec(stateQyery.CreateDefaultStates, UserID)
	if err != nil {
		return err
	}

	return nil
}
