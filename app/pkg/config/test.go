package config

var test = Config{
	Database: DatabaseConfig{
		User:                 "root",
		Password:             "",
		Host:                 "127.0.0.1",
		Port:                 "3306",
		Database:             "hakushi_test",
		TimeZone:             "Asia/Tokyo",
		AllowNativePasswords: true,
		ParseTime:            true,
	},
	Server: ServerConfig{
		Host: "0.0.0.0",
		Port: "80",
	},
}
