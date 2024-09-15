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
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}
