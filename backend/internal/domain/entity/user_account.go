package entity

import (
	"time"
)

type UserAccount struct {
	ID        int64
	Sub       string
	Username  string
	Email     string
	Name      string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserAccount(sub, username, email, name, imageURL string) *UserAccount {
	now := time.Now()
	return &UserAccount{
		Sub:       sub,
		Username:  username,
		Email:     email,
		Name:      name,
		ImageURL:  imageURL,
		CreatedAt: now,
		UpdatedAt: now,
	}
}