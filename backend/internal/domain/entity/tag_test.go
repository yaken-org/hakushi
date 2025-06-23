package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	name := "test-tag"

	tag := NewTag(name)

	assert.Equal(t, name, tag.Name)
	assert.False(t, tag.CreatedAt.IsZero())
	assert.False(t, tag.UpdatedAt.IsZero())
	assert.Equal(t, tag.CreatedAt, tag.UpdatedAt)
}

func TestNewTag_EmptyName(t *testing.T) {
	tag := NewTag("")

	assert.Empty(t, tag.Name)
	assert.False(t, tag.CreatedAt.IsZero())
	assert.False(t, tag.UpdatedAt.IsZero())
}

func TestNewPostTag(t *testing.T) {
	postID := int64(1)
	tagID := int64(10)

	postTag := NewPostTag(postID, tagID)

	assert.Equal(t, postID, postTag.PostID)
	assert.Equal(t, tagID, postTag.TagID)
	assert.False(t, postTag.CreatedAt.IsZero())
}