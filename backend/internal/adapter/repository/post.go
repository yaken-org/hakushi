package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindByID(ctx context.Context, id int64) (*entity.Post, error) {
	query := "SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	
	post := &entity.Post{}
	err := row.Scan(
		&post.ID,
		&post.UserAccountID,
		&post.ImageID,
		&post.Title,
		&post.Content,
		&post.Likes,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPostNotFound
		}
		return nil, err
	}
	
	return post, nil
}

func (r *postRepository) FindAll(ctx context.Context) ([]*entity.Post, error) {
	query := "SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post ORDER BY created_at DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(
			&post.ID,
			&post.UserAccountID,
			&post.ImageID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

func (r *postRepository) FindByUserAccountID(ctx context.Context, userAccountID int64) ([]*entity.Post, error) {
	query := "SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE user_account_id = ? ORDER BY created_at DESC"
	rows, err := r.db.QueryContext(ctx, query, userAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(
			&post.ID,
			&post.UserAccountID,
			&post.ImageID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

func (r *postRepository) FindByTitleContaining(ctx context.Context, title string) ([]*entity.Post, error) {
	query := "SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE title LIKE ? ORDER BY created_at DESC"
	rows, err := r.db.QueryContext(ctx, query, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(
			&post.ID,
			&post.UserAccountID,
			&post.ImageID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

func (r *postRepository) FindOrderByLikes(ctx context.Context) ([]*entity.Post, error) {
	query := "SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE likes > 0 ORDER BY likes DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*entity.Post
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(
			&post.ID,
			&post.UserAccountID,
			&post.ImageID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

func (r *postRepository) Create(ctx context.Context, post *entity.Post) error {
	query := "INSERT INTO post (user_account_id, image_id, title, content, likes, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query,
		post.UserAccountID,
		post.ImageID,
		post.Title,
		post.Content,
		post.Likes,
		post.CreatedAt,
		post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	post.ID = id
	return nil
}

func (r *postRepository) Update(ctx context.Context, post *entity.Post) error {
	post.UpdatedAt = time.Now()
	query := "UPDATE post SET title = ?, content = ?, likes = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query,
		post.Title,
		post.Content,
		post.Likes,
		post.UpdatedAt,
		post.ID,
	)
	return err
}

func (r *postRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM post WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}