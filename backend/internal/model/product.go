package model

import (
	"time"
)

type Product struct {
	ID int64 `json:"id"`

	Name string `json:"name"`
	Link string `json:"link"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) FromRow(row Scannable) error {
	return row.Scan(
		&p.ID,
		&p.Name,
		&p.Link,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
}
