package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID int64 `json:"id"`

	Name string `json:"name"`
	Link string `json:"link"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) FromRow(row *sql.Row) error {
	return row.Scan(
		&p.ID,
		&p.Name,
		&p.Link,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}