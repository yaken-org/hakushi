package database

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yaken-org/hakushi/pkg/config"
)

// dsn はデータベース接続情報を返す
func dsn(c config.Config) string {
	tz, err := time.LoadLocation(c.Database.TimeZone)
	if err != nil {
		// デフォルトはUTC
		tz = time.UTC
	}

	cfg := mysql.Config{
		User:                 c.Database.User,
		Passwd:               c.Database.Password,
		Net:                  "tcp",
		Addr:                 c.Database.Host + ":" + c.Database.Port,
		DBName:               c.Database.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  tz,
	}

	return cfg.FormatDSN()
}
