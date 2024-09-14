package model

import (
	"database/sql"
	"time"
)

type PostItem struct {
	ID int64 `json:"id"`

	PostID    int64 `json:"post_id"`
	ProductID int64 `json:"product_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *PostItem) FromRow(row sql.Row) error {
	return row.Scan(
		&p.ID,
		&p.PostID,
		&p.ProductID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}
