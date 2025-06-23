package usecase

import (
	"context"
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrConflict         = errors.New("conflict")
	ErrInternalServer   = errors.New("internal server error")
)

type Transaction interface {
	Begin(ctx context.Context) (Transaction, error)
	Commit() error
	Rollback() error
}

type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}