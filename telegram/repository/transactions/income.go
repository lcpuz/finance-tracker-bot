package transactionsRepository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	transactionQuery "github.com/lcpuz/finance-tracker-bot/telegram/repository/query/transaction"
	"github.com/lcpuz/finance-tracker-bot/telegram/request"
)

type IncomePostgres struct {
	db *sqlx.DB
}

func NewIncomePostgres(db *sqlx.DB) *IncomePostgres {
	return &IncomePostgres{db: db}
}

func (r *IncomePostgres) AddUncategorizedIncome(UncategorizedIncome request.AddUncategorizedIncomeRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, transactionQuery.AddUncategorizedIncome,
		UncategorizedIncome.UserID,
		UncategorizedIncome.TransactionAmount,
		UncategorizedIncome.TransactionDescription)

	if err != nil {
		return err
	}

	return nil
}
