package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yaken-org/hakushi/internal/config"
)

func New(config *config.Config) (*sql.DB, error) {
	timezone, err := time.LoadLocation(config.Database.Timezone)
	if err != nil {
		return nil, err
	}

	dsn := &mysql.Config{
		User:                 config.Database.Username,
		Passwd:               config.Database.Password,
		Net:                  "tcp",
		Addr:                 config.Database.Host,
		DBName:               config.Database.Database,
		ParseTime:            true,
		Loc:                  timezone,
		AllowNativePasswords: true,
	}

	database, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return nil, err
	}

	database.SetMaxIdleConns(config.Database.MaxIdleConns)
	database.SetMaxOpenConns(config.Database.MaxOpenConns)
	database.SetConnMaxLifetime(time.Duration(config.Database.ConnMaxLifetime) * time.Second)

	return database, nil
}