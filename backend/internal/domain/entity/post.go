package entity

import (
	"time"
)

type Post struct {
	ID            int64
	UserAccountID int64
	ImageID       int64
	Title         string
	Content       string
	Likes         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPost(userAccountID, imageID int64, title, content string) *Post {
	now := time.Now()
	return &Post{
		UserAccountID: userAccountID,
		ImageID:       imageID,
		Title:         title,
		Content:       content,
		Likes:         0,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func (p *Post) IncrementLikes() {
	p.Likes++
	p.UpdatedAt = time.Now()
}