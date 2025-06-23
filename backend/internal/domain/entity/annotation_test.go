package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnnotation(t *testing.T) {
	postID := int64(1)
	productID := int64(100)
	displayName := "Test Product"
	x := 150
	y := 300

	annotation := NewAnnotation(postID, &productID, displayName, x, y)

	assert.Equal(t, postID, annotation.PostID)
	assert.Equal(t, &productID, annotation.ProductID)
	assert.Equal(t, displayName, annotation.DisplayName)
	assert.Equal(t, x, annotation.X)
	assert.Equal(t, y, annotation.Y)
	assert.False(t, annotation.CreatedAt.IsZero())
	assert.False(t, annotation.UpdatedAt.IsZero())
	assert.Equal(t, annotation.CreatedAt, annotation.UpdatedAt)
}

func TestNewAnnotation_NilProductID(t *testing.T) {
	postID := int64(1)
	displayName := "Test Annotation"
	x := 150
	y := 300

	annotation := NewAnnotation(postID, nil, displayName, x, y)

	assert.Equal(t, postID, annotation.PostID)
	assert.Nil(t, annotation.ProductID)
	assert.Equal(t, displayName, annotation.DisplayName)
	assert.Equal(t, x, annotation.X)
	assert.Equal(t, y, annotation.Y)
}

func TestNewAnnotation_ZeroCoordinates(t *testing.T) {
	annotation := NewAnnotation(1, nil, "Test", 0, 0)

	assert.Equal(t, 0, annotation.X)
	assert.Equal(t, 0, annotation.Y)
}