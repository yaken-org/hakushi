package database

import (
	"database/sql"

	"github.com/yaken-org/hakushi/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var database *Database

func New() *Database {
	return database
}

func init() {
	cfg := config.Get()
	db, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		panic(err)
	}

	g, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database = &Database{DB: g}
}
