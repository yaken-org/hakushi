package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/testutil/helper"
)

func TestPostRepository_FindByID(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		postID := int64(1)
		userAccountID := int64(1)
		imageID := int64(100)
		title := "Test Post"
		content := "Test Content"
		likes := 5

		helper.ExpectFindPostByID(mock, postID, userAccountID, imageID, title, content, likes)

		post, err := repo.FindByID(ctx, postID)

		assert.NoError(t, err)
		assert.Equal(t, postID, post.ID)
		assert.Equal(t, userAccountID, post.UserAccountID)
		assert.Equal(t, imageID, post.ImageID)
		assert.Equal(t, title, post.Title)
		assert.Equal(t, content, post.Content)
		assert.Equal(t, likes, post.Likes)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		postID := int64(999)

		mock.ExpectQuery("SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE id = ?").
			WithArgs(postID).
			WillReturnError(sql.ErrNoRows)

		post, err := repo.FindByID(ctx, postID)

		assert.Error(t, err)
		assert.Equal(t, ErrPostNotFound, err)
		assert.Nil(t, post)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestPostRepository_FindAll(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		posts := [][]interface{}{
			{int64(1), int64(1), int64(100), "Post 1", "Content 1", 5, time.Now(), time.Now()},
			{int64(2), int64(1), int64(101), "Post 2", "Content 2", 3, time.Now(), time.Now()},
		}

		helper.ExpectFindAllPosts(mock, posts)

		result, err := repo.FindAll(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(1), result[0].ID)
		assert.Equal(t, "Post 1", result[0].Title)
		assert.Equal(t, int64(2), result[1].ID)
		assert.Equal(t, "Post 2", result[1].Title)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty result", func(t *testing.T) {
		helper.ExpectFindAllPosts(mock, [][]interface{}{})

		result, err := repo.FindAll(ctx)

		assert.NoError(t, err)
		assert.Empty(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestPostRepository_Create(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		post := entity.NewPost(1, 100, "Test Post", "Test Content")
		insertID := int64(123)

		helper.ExpectCreatePost(mock, post.UserAccountID, post.ImageID, post.Title, post.Content, insertID)

		err := repo.Create(ctx, post)

		assert.NoError(t, err)
		assert.Equal(t, insertID, post.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestPostRepository_Update(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		post := &entity.Post{
			ID:            1,
			UserAccountID: 1,
			ImageID:       100,
			Title:         "Updated Title",
			Content:       "Updated Content",
			Likes:         10,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		helper.ExpectUpdatePost(mock, post.ID, post.Title, post.Content, post.Likes)

		err := repo.Update(ctx, post)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestPostRepository_Delete(t *testing.T) {
	mock, db, err := helper.SetupMockDB()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostRepository(db)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		postID := int64(1)

		helper.ExpectDeletePost(mock, postID)

		err := repo.Delete(ctx, postID)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}