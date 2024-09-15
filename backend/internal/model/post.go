package model

import (
	"time"
)

type Post struct {
	ID int64 `json:"id"`

	UserAccountID int64 `json:"user_account_id"`
	ImageID       int64 `json:"image_id"`

	Title   string `json:"title"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Post) FromRow(row Scannable) error {
	return row.Scan(
		&p.ID,
		&p.UserAccountID,
		&p.ImageID,
		&p.Title,
		&p.Content,
		&p.Likes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}

type APIPost struct {
	ID int64 `json:"id"`

	UserAccountID int64 `json:"user_account_id"`
	ImageID       int64 `json:"image_id"`

	Title   string `json:"title"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`

	Annotations []*Annotation `json:"annotations"`
	Tags        []*Tag        `json:"tags"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Post) ToAPIPost(annotations []*Annotation, tags []*Tag) *APIPost {
	if annotations == nil {
		annotations = []*Annotation{}
	}
	if tags == nil {
		tags = []*Tag{}
	}
	return &APIPost{
		ID:            p.ID,
		UserAccountID: p.UserAccountID,
		ImageID:       p.ImageID,
		Title:         p.Title,
		Content:       p.Content,
		Likes:         p.Likes,
		Annotations:   annotations,
		Tags:          tags,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
