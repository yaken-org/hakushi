package repository

import (
	"context"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.Product, error)
	FindAll(ctx context.Context) ([]*entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id int64) error
}