package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yaken-org/hakushi/internal/config"
	gormMySQL "gorm.io/driver/mysql"
)

type Database struct {
	*sql.DB
	Gorm *gorm.DB
}

var db *Database

func New() *Database {
	if db == nil {
		panic("database is not initialized")
	}

	return db
}

func Initialize(config *config.Config) error {
	if db != nil {
		return nil
	}

	timezone, err := time.LoadLocation(config.Database.Timezone)
	if err != nil {
		return err
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
		return err
	}

	database.SetMaxIdleConns(config.Database.MaxIdleConns)
	database.SetMaxOpenConns(config.Database.MaxOpenConns)
	database.SetConnMaxLifetime(time.Duration(config.Database.ConnMaxLifetime) * time.Second)

	var gormDB *gorm.DB
	if gormDB, err = gorm.Open(gormMySQL.New(gormMySQL.Config{
		Conn: database,
	}), &gorm.Config{}); err != nil {
		panic(err)
	}

	db = &Database{
		DB:   database,
		Gorm: gormDB,
	}
	return nil
}
