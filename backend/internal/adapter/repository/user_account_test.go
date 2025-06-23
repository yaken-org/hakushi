package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/testutil/helper"
)

func TestUserAccountRepository_FindByID(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserAccountRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userID := int64(1)
		sub := "google-oauth2|123456"
		username := "testuser"
		email := "test@example.com"
		name := "Test User"
		imageURL := "https://example.com/avatar.jpg"
		now := time.Now()

		rows := sqlmock.NewRows([]string{
			"id", "sub", "username", "email", "name", "image_url", "created_at", "updated_at",
		}).AddRow(userID, sub, username, email, name, imageURL, now, now)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE id = ?")).
			WithArgs(userID).
			WillReturnRows(rows)

		user, err := repo.FindByID(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, sub, user.Sub)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, imageURL, user.ImageURL)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		userID := int64(999)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE id = ?")).
			WithArgs(userID).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.FindByID(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, ErrUserAccountNotFound, err)
		assert.Nil(t, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserAccountRepository_FindBySub(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserAccountRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		sub := "google-oauth2|123456"
		userID := int64(1)
		username := "testuser"
		email := "test@example.com"
		name := "Test User"
		imageURL := "https://example.com/avatar.jpg"
		now := time.Now()

		rows := sqlmock.NewRows([]string{
			"id", "sub", "username", "email", "name", "image_url", "created_at", "updated_at",
		}).AddRow(userID, sub, username, email, name, imageURL, now, now)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, sub, username, email, name, image_url, created_at, updated_at FROM user_account WHERE sub = ?")).
			WithArgs(sub).
			WillReturnRows(rows)

		user, err := repo.FindBySub(ctx, sub)

		assert.NoError(t, err)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, sub, user.Sub)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserAccountRepository_Create(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserAccountRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userAccount := entity.NewUserAccount(
			"google-oauth2|123456",
			"testuser",
			"test@example.com",
			"Test User",
			"https://example.com/avatar.jpg",
		)
		insertID := int64(123)

		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO user_account (sub, username, email, name, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")).
			WithArgs(
				userAccount.Sub,
				userAccount.Username,
				userAccount.Email,
				userAccount.Name,
				userAccount.ImageURL,
				helper.AnyTime{},
				helper.AnyTime{},
			).
			WillReturnResult(sqlmock.NewResult(insertID, 1))

		err := repo.Create(ctx, userAccount)

		assert.NoError(t, err)
		assert.Equal(t, insertID, userAccount.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserAccountRepository_Update(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserAccountRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userAccount := &entity.UserAccount{
			ID:        1,
			Sub:       "google-oauth2|123456",
			Username:  "updateduser",
			Email:     "updated@example.com",
			Name:      "Updated User",
			ImageURL:  "https://example.com/new-avatar.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mock.ExpectExec(regexp.QuoteMeta("UPDATE user_account SET username = ?, email = ?, name = ?, image_url = ?, updated_at = ? WHERE id = ?")).
			WithArgs(
				userAccount.Username,
				userAccount.Email,
				userAccount.Name,
				userAccount.ImageURL,
				helper.AnyTime{},
				userAccount.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(ctx, userAccount)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}