package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserAccount(t *testing.T) {
	sub := "google-oauth2|123456"
	username := "testuser"
	email := "test@example.com"
	name := "Test User"
	imageURL := "https://example.com/avatar.jpg"

	userAccount := NewUserAccount(sub, username, email, name, imageURL)

	assert.Equal(t, sub, userAccount.Sub)
	assert.Equal(t, username, userAccount.Username)
	assert.Equal(t, email, userAccount.Email)
	assert.Equal(t, name, userAccount.Name)
	assert.Equal(t, imageURL, userAccount.ImageURL)
	assert.False(t, userAccount.CreatedAt.IsZero())
	assert.False(t, userAccount.UpdatedAt.IsZero())
	assert.Equal(t, userAccount.CreatedAt, userAccount.UpdatedAt)
}

func TestNewUserAccount_EmptyValues(t *testing.T) {
	userAccount := NewUserAccount("", "", "", "", "")

	assert.Empty(t, userAccount.Sub)
	assert.Empty(t, userAccount.Username)
	assert.Empty(t, userAccount.Email)
	assert.Empty(t, userAccount.Name)
	assert.Empty(t, userAccount.ImageURL)
	assert.False(t, userAccount.CreatedAt.IsZero())
	assert.False(t, userAccount.UpdatedAt.IsZero())
}