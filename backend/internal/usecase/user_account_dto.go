package usecase

import (
	"time"
)

type CreateUserAccountInput struct {
	Sub      string
	Username string
	Email    string
	Name     string
	ImageURL string
}

type UserAccountOutput struct {
	ID        int64     `json:"id"`
	Sub       string    `json:"sub"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserAccountInput struct {
	ID       int64
	Username string
	Email    string
	Name     string
	ImageURL string
}