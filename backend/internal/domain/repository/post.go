package repository

import (
	"context"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type PostRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.Post, error)
	FindAll(ctx context.Context) ([]*entity.Post, error)
	FindByUserAccountID(ctx context.Context, userAccountID int64) ([]*entity.Post, error)
	FindByTitleContaining(ctx context.Context, title string) ([]*entity.Post, error)
	FindOrderByLikes(ctx context.Context) ([]*entity.Post, error)
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id int64) error
}