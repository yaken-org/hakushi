package model

import "database/sql"

type Model interface {
	FromRow(row sql.Row) error
}
