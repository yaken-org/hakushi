package model

import "database/sql"

type Model interface {
	FromRow(row Scannable) error
}

type Scannable interface {
	Scan(dest ...any) error
}

func QueryRow[T Model](db *sql.DB, model T, query string, args ...any) error {
	row := db.QueryRow(query, args...)
	return model.FromRow(row)
}
