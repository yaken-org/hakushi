package database

import (
	"context"
	"database/sql"

	"github.com/yaken-org/hakushi/internal/usecase"
)

type transactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) usecase.TransactionManager {
	return &transactionManager{db: db}
}

func (tm *transactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	ctxWithTx := context.WithValue(ctx, "tx", tx)
	
	if err := fn(ctxWithTx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}