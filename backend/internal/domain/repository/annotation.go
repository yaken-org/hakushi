package repository

import (
	"context"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type AnnotationRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.Annotation, error)
	FindByPostID(ctx context.Context, postID int64) ([]*entity.Annotation, error)
	FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Annotation, error)
	Create(ctx context.Context, annotation *entity.Annotation) error
	Update(ctx context.Context, annotation *entity.Annotation) error
	Delete(ctx context.Context, id int64) error
}