package model

import "time"

type Annotation struct {
	ID int64 `json:"id"`

	PostID    int64 `json:"post_id"`
	ProductID int64 `json:"product_id"`

	DisplayName string `json:"display_name"`

	X int64 `json:"x"`
	Y int64 `json:"y"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Annotation) FromRow(row Scannable) error {
	return row.Scan(
		&a.ID,
		&a.PostID,
		&a.ProductID,
		&a.DisplayName,
		&a.X,
		&a.Y,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
}
