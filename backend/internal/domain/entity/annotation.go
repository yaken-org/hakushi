package entity

import (
	"time"
)

type Annotation struct {
	ID          int64
	PostID      int64
	ProductID   *int64
	DisplayName string
	X           int
	Y           int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAnnotation(postID int64, productID *int64, displayName string, x, y int) *Annotation {
	now := time.Now()
	return &Annotation{
		PostID:      postID,
		ProductID:   productID,
		DisplayName: displayName,
		X:           x,
		Y:           y,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}