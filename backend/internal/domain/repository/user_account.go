package repository

import (
	"context"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type UserAccountRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.UserAccount, error)
	FindBySub(ctx context.Context, sub string) (*entity.UserAccount, error)
	FindByUsername(ctx context.Context, username string) (*entity.UserAccount, error)
	Create(ctx context.Context, userAccount *entity.UserAccount) error
	Update(ctx context.Context, userAccount *entity.UserAccount) error
	Delete(ctx context.Context, id int64) error
}