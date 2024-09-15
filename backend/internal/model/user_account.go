package model

import (
	"time"
)

type UserAccount struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserAccount) FromRow(row Scannable) error {
	return row.Scan(
		&u.ID,
		&u.Name,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}
