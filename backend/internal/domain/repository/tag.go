package repository

import (
	"context"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type TagRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.Tag, error)
	FindByName(ctx context.Context, name string) (*entity.Tag, error)
	FindAll(ctx context.Context) ([]*entity.Tag, error)
	FindByPostID(ctx context.Context, postID int64) ([]*entity.Tag, error)
	FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Tag, error)
	Create(ctx context.Context, tag *entity.Tag) error
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}

type PostTagRepository interface {
	FindByPostID(ctx context.Context, postID int64) ([]*entity.PostTag, error)
	FindByTagID(ctx context.Context, tagID int64) ([]*entity.PostTag, error)
	Create(ctx context.Context, postTag *entity.PostTag) error
	Delete(ctx context.Context, postID, tagID int64) error
}