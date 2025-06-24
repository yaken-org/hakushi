package usecase

import (
	"time"
)

type CreatePostInput struct {
	UserAccountID int64
	ImageID       int64
	Title         string
	Content       string
	Annotations   []CreateAnnotationInput
	Tags          []string
}

type CreateAnnotationInput struct {
	ProductID   *int64
	DisplayName string
	X           int
	Y           int
}

type PostOutput struct {
	ID            int64                `json:"id"`
	UserAccountID int64                `json:"user_account_id"`
	ImageID       int64                `json:"image_id"`
	Title         string               `json:"title"`
	Content       string               `json:"content"`
	Likes         int                  `json:"likes"`
	Annotations   []AnnotationOutput   `json:"annotations"`
	Tags          []TagOutput          `json:"tags"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

type AnnotationOutput struct {
	ID          int64  `json:"id"`
	PostID      int64  `json:"post_id"`
	ProductID   *int64 `json:"product_id"`
	DisplayName string `json:"display_name"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TagOutput struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePostInput struct {
	ID      int64
	Title   string
	Content string
}