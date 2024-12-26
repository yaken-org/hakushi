package config

import (
	"os"
	"strconv"
)

var env = map[string]Config{
	"test":        test,
	"local":       local,
	"development": development,
	"production":  production,
}

var cfg Config

func SetEnv(e string, c Config) {
	env[e] = c
}

func Get() Config {
	if cfg == (Config{}) {
		setup()
	}
	return cfg
}

func setup() {
	e := os.Getenv("APP_ENV")
	if e == "" {
		e = "local"
	}
	c := env[e]

	// 環境変数が設定されている場合は上書き
	c.Database = DatabaseConfig{
		User:                 load("DB_USER", c.Database.User),
		Password:             load("DB_PASSWORD", c.Database.Password),
		Host:                 load("DB_HOST", c.Database.Host),
		Port:                 load("DB_PORT", c.Database.Port),
		Database:             load("DB_DATABASE", c.Database.Database),
		TimeZone:             load("DB_TIMEZONE", c.Database.TimeZone),
		AllowNativePasswords: parseBool(load("DB_ALLOW_NATIVE_PASSWORDS", c.Database.AllowNativePasswords)),
		ParseTime:            parseBool(load("DB_PARSE_TIME", c.Database.ParseTime)),
	}
	c.Server = ServerConfig{
		Host: load("SERVER_HOST", c.Server.Host),
		Port: load("SERVER_PORT", c.Server.Port),
	}

	cfg = c
}

// parseBool は文字列をbool型に変換する
func parseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

// fallbackValue はload関数の引数の型を制限するためのinterface
type fallbackValue interface {
	~int | ~string | ~bool
}

// load は環境変数が設定されている場合はその値を返し、設定されていない場合はfallbacksの値を返す
// fallbacksの値は優先順位が高い順に設定し、空文字の場合は次の値を評価する。
func load[T fallbackValue](key string, fallbacks ...T) string {
	v := os.Getenv(key)
	if v == "" {
		for _, f := range fallbacks {
			switch fv := any(f).(type) {
			case string:
				if fv != "" {
					v = fv
					break
				}
			case int:
				v = strconv.Itoa(fv)
				break
			case bool:
				v = strconv.FormatBool(fv)
				break
			}
		}
	}

	return v
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	User                 string
	Password             string
	Host                 string
	Port                 string
	Database             string
	TimeZone             string
	AllowNativePasswords bool
	ParseTime            bool
}

type ServerConfig struct {
	Host string
	Port string
}
