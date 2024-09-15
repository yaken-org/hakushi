package model

import "time"

type Tag struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tag) FromRow(row Scannable) error {
	return row.Scan(
		&t.ID,
		&t.Name,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
}
