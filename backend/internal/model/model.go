package model

import "database/sql"

type Model interface {
	FromRow(row *sql.Row) error
}

func QueryRow[T Model](db *sql.DB, model T, query string, args ...any) (T, error) {
	row := db.QueryRow(query, args...)
	err := model.FromRow(row)
	return model, err
}
