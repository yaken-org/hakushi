package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPost(t *testing.T) {
	userAccountID := int64(1)
	imageID := int64(100)
	title := "Test Post"
	content := "This is a test post"

	post := NewPost(userAccountID, imageID, title, content)

	assert.Equal(t, userAccountID, post.UserAccountID)
	assert.Equal(t, imageID, post.ImageID)
	assert.Equal(t, title, post.Title)
	assert.Equal(t, content, post.Content)
	assert.Equal(t, 0, post.Likes)
	assert.False(t, post.CreatedAt.IsZero())
	assert.False(t, post.UpdatedAt.IsZero())
	assert.Equal(t, post.CreatedAt, post.UpdatedAt)
}

func TestPost_IncrementLikes(t *testing.T) {
	post := NewPost(1, 100, "Test", "Content")
	originalLikes := post.Likes
	originalUpdatedAt := post.UpdatedAt

	time.Sleep(1 * time.Millisecond)
	post.IncrementLikes()

	assert.Equal(t, originalLikes+1, post.Likes)
	assert.True(t, post.UpdatedAt.After(originalUpdatedAt))
}

func TestPost_IncrementLikes_Multiple(t *testing.T) {
	post := NewPost(1, 100, "Test", "Content")

	for i := 0; i < 5; i++ {
		post.IncrementLikes()
	}

	assert.Equal(t, 5, post.Likes)
}