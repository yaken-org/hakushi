package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type userAccountRepository struct {
	db *sql.DB
}

func NewUserAccountRepository(db *sql.DB) repository.UserAccountRepository {
	return &userAccountRepository{db: db}
}

func (r *userAccountRepository) FindByID(ctx context.Context, id int64) (*entity.UserAccount, error) {
	query := "SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	
	userAccount := &entity.UserAccount{}
	err := row.Scan(
		&userAccount.ID,
		&userAccount.Sub,
		&userAccount.Username,
		&userAccount.Email,
		&userAccount.Name,
		&userAccount.ImageURL,
		&userAccount.CreatedAt,
		&userAccount.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserAccountNotFound
		}
		return nil, err
	}
	
	return userAccount, nil
}

func (r *userAccountRepository) FindBySub(ctx context.Context, sub string) (*entity.UserAccount, error) {
	query := "SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE sub = ?"
	row := r.db.QueryRowContext(ctx, query, sub)
	
	userAccount := &entity.UserAccount{}
	err := row.Scan(
		&userAccount.ID,
		&userAccount.Sub,
		&userAccount.Username,
		&userAccount.Email,
		&userAccount.Name,
		&userAccount.ImageURL,
		&userAccount.CreatedAt,
		&userAccount.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserAccountNotFound
		}
		return nil, err
	}
	
	return userAccount, nil
}

func (r *userAccountRepository) FindByUsername(ctx context.Context, username string) (*entity.UserAccount, error) {
	query := "SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE username = ?"
	row := r.db.QueryRowContext(ctx, query, username)
	
	userAccount := &entity.UserAccount{}
	err := row.Scan(
		&userAccount.ID,
		&userAccount.Sub,
		&userAccount.Username,
		&userAccount.Email,
		&userAccount.Name,
		&userAccount.ImageURL,
		&userAccount.CreatedAt,
		&userAccount.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserAccountNotFound
		}
		return nil, err
	}
	
	return userAccount, nil
}

func (r *userAccountRepository) Create(ctx context.Context, userAccount *entity.UserAccount) error {
	query := "INSERT INTO user_account (sub, username, email, name, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query,
		userAccount.Sub,
		userAccount.Username,
		userAccount.Email,
		userAccount.Name,
		userAccount.ImageURL,
		userAccount.CreatedAt,
		userAccount.UpdatedAt,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	userAccount.ID = id
	return nil
}

func (r *userAccountRepository) Update(ctx context.Context, userAccount *entity.UserAccount) error {
	userAccount.UpdatedAt = time.Now()
	query := "UPDATE user_account SET username = ?, email = ?, name = ?, image_url = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query,
		userAccount.Username,
		userAccount.Email,
		userAccount.Name,
		userAccount.ImageURL,
		userAccount.UpdatedAt,
		userAccount.ID,
	)
	return err
}

func (r *userAccountRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM user_account WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}