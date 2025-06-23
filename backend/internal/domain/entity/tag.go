package entity

import (
	"time"
)

type Tag struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTag(name string) *Tag {
	now := time.Now()
	return &Tag{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type PostTag struct {
	ID        int64
	PostID    int64
	TagID     int64
	CreatedAt time.Time
}

func NewPostTag(postID, tagID int64) *PostTag {
	return &PostTag{
		PostID:    postID,
		TagID:     tagID,
		CreatedAt: time.Now(),
	}
}