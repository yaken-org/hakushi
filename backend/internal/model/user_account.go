package model

import (
	"time"
)

type UserAccount struct {
	ID int64 `json:"id"`

	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	IconURL     string `json:"icon_url"`

	Sub string `json:"sub"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserAccount) FromRow(row Scannable) error {
	return row.Scan(
		&u.ID,

		&u.Name,
		&u.DisplayName,
		&u.IconURL,

		&u.Sub,

		&u.CreatedAt,
		&u.UpdatedAt,
	)
}
