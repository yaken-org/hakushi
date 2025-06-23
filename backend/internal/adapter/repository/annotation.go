package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type annotationRepository struct {
	db *sql.DB
}

func NewAnnotationRepository(db *sql.DB) repository.AnnotationRepository {
	return &annotationRepository{db: db}
}

func (r *annotationRepository) FindByID(ctx context.Context, id int64) (*entity.Annotation, error) {
	query := "SELECT id, post_id, product_id, display_name, x, y, created_at, updated_at FROM annotation WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	
	annotation := &entity.Annotation{}
	err := row.Scan(
		&annotation.ID,
		&annotation.PostID,
		&annotation.ProductID,
		&annotation.DisplayName,
		&annotation.X,
		&annotation.Y,
		&annotation.CreatedAt,
		&annotation.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAnnotationNotFound
		}
		return nil, err
	}
	
	return annotation, nil
}

func (r *annotationRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.Annotation, error) {
	query := "SELECT id, post_id, product_id, display_name, x, y, created_at, updated_at FROM annotation WHERE post_id = ? ORDER BY created_at"
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var annotations []*entity.Annotation
	for rows.Next() {
		annotation := &entity.Annotation{}
		err := rows.Scan(
			&annotation.ID,
			&annotation.PostID,
			&annotation.ProductID,
			&annotation.DisplayName,
			&annotation.X,
			&annotation.Y,
			&annotation.CreatedAt,
			&annotation.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		annotations = append(annotations, annotation)
	}
	
	return annotations, nil
}

func (r *annotationRepository) FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Annotation, error) {
	if len(postIDs) == 0 {
		return make(map[int64][]*entity.Annotation), nil
	}
	
	placeholders := make([]string, len(postIDs))
	args := make([]interface{}, len(postIDs))
	for i, id := range postIDs {
		placeholders[i] = "?"
		args[i] = id
	}
	
	query := fmt.Sprintf("SELECT id, post_id, product_id, display_name, x, y, created_at, updated_at FROM annotation WHERE post_id IN (%s) ORDER BY post_id, created_at", strings.Join(placeholders, ","))
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	result := make(map[int64][]*entity.Annotation)
	for rows.Next() {
		annotation := &entity.Annotation{}
		err := rows.Scan(
			&annotation.ID,
			&annotation.PostID,
			&annotation.ProductID,
			&annotation.DisplayName,
			&annotation.X,
			&annotation.Y,
			&annotation.CreatedAt,
			&annotation.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result[annotation.PostID] = append(result[annotation.PostID], annotation)
	}
	
	return result, nil
}

func (r *annotationRepository) Create(ctx context.Context, annotation *entity.Annotation) error {
	query := "INSERT INTO annotation (post_id, product_id, display_name, x, y, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query,
		annotation.PostID,
		annotation.ProductID,
		annotation.DisplayName,
		annotation.X,
		annotation.Y,
		annotation.CreatedAt,
		annotation.UpdatedAt,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	annotation.ID = id
	return nil
}

func (r *annotationRepository) Update(ctx context.Context, annotation *entity.Annotation) error {
	query := "UPDATE annotation SET product_id = ?, display_name = ?, x = ?, y = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query,
		annotation.ProductID,
		annotation.DisplayName,
		annotation.X,
		annotation.Y,
		annotation.UpdatedAt,
		annotation.ID,
	)
	return err
}

func (r *annotationRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM annotation WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}