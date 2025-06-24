package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) repository.TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) FindByID(ctx context.Context, id int64) (*entity.Tag, error) {
	query := "SELECT id, name, created_at, updated_at FROM tag WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	
	tag := &entity.Tag{}
	err := row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTagNotFound
		}
		return nil, err
	}
	
	return tag, nil
}

func (r *tagRepository) FindByName(ctx context.Context, name string) (*entity.Tag, error) {
	query := "SELECT id, name, created_at, updated_at FROM tag WHERE name = ?"
	row := r.db.QueryRowContext(ctx, query, name)
	
	tag := &entity.Tag{}
	err := row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTagNotFound
		}
		return nil, err
	}
	
	return tag, nil
}

func (r *tagRepository) FindAll(ctx context.Context) ([]*entity.Tag, error) {
	query := "SELECT id, name, created_at, updated_at FROM tag ORDER BY name"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var tags []*entity.Tag
	for rows.Next() {
		tag := &entity.Tag{}
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	
	return tags, nil
}

func (r *tagRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.Tag, error) {
	query := `
		SELECT t.id, t.name, t.created_at, t.updated_at 
		FROM tag t 
		INNER JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id = ? 
		ORDER BY t.name
	`
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var tags []*entity.Tag
	for rows.Next() {
		tag := &entity.Tag{}
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	
	return tags, nil
}

func (r *tagRepository) FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Tag, error) {
	if len(postIDs) == 0 {
		return make(map[int64][]*entity.Tag), nil
	}
	
	placeholders := make([]string, len(postIDs))
	args := make([]interface{}, len(postIDs))
	for i, id := range postIDs {
		placeholders[i] = "?"
		args[i] = id
	}
	
	query := fmt.Sprintf(`
		SELECT t.id, t.name, t.created_at, t.updated_at, pt.post_id
		FROM tag t 
		INNER JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id IN (%s)
		ORDER BY pt.post_id, t.name
	`, strings.Join(placeholders, ","))
	
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	result := make(map[int64][]*entity.Tag)
	for rows.Next() {
		tag := &entity.Tag{}
		var postID int64
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&postID,
		)
		if err != nil {
			return nil, err
		}
		result[postID] = append(result[postID], tag)
	}
	
	return result, nil
}

func (r *tagRepository) Create(ctx context.Context, tag *entity.Tag) error {
	query := "INSERT INTO tag (name, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query,
		tag.Name,
		tag.CreatedAt,
		tag.UpdatedAt,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	tag.ID = id
	return nil
}

func (r *tagRepository) Update(ctx context.Context, tag *entity.Tag) error {
	query := "UPDATE tag SET name = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query,
		tag.Name,
		tag.UpdatedAt,
		tag.ID,
	)
	return err
}

func (r *tagRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM tag WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

type postTagRepository struct {
	db *sql.DB
}

func NewPostTagRepository(db *sql.DB) repository.PostTagRepository {
	return &postTagRepository{db: db}
}

func (r *postTagRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.PostTag, error) {
	query := "SELECT id, post_id, tag_id, created_at FROM post_tag WHERE post_id = ?"
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var postTags []*entity.PostTag
	for rows.Next() {
		postTag := &entity.PostTag{}
		err := rows.Scan(
			&postTag.ID,
			&postTag.PostID,
			&postTag.TagID,
			&postTag.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		postTags = append(postTags, postTag)
	}
	
	return postTags, nil
}

func (r *postTagRepository) FindByTagID(ctx context.Context, tagID int64) ([]*entity.PostTag, error) {
	query := "SELECT id, post_id, tag_id, created_at FROM post_tag WHERE tag_id = ?"
	rows, err := r.db.QueryContext(ctx, query, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var postTags []*entity.PostTag
	for rows.Next() {
		postTag := &entity.PostTag{}
		err := rows.Scan(
			&postTag.ID,
			&postTag.PostID,
			&postTag.TagID,
			&postTag.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		postTags = append(postTags, postTag)
	}
	
	return postTags, nil
}

func (r *postTagRepository) Create(ctx context.Context, postTag *entity.PostTag) error {
	query := "INSERT INTO post_tag (post_id, tag_id, created_at) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query,
		postTag.PostID,
		postTag.TagID,
		postTag.CreatedAt,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	postTag.ID = id
	return nil
}

func (r *postTagRepository) Delete(ctx context.Context, postID, tagID int64) error {
	query := "DELETE FROM post_tag WHERE post_id = ? AND tag_id = ?"
	_, err := r.db.ExecContext(ctx, query, postID, tagID)
	return err
}