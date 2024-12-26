package config

var local = Config{
	Database: DatabaseConfig{
		User:                 "root",
		Password:             "",
		Host:                 "db",
		Port:                 "3306",
		Database:             "hakushi",
		TimeZone:             "Asia/Tokyo",
		AllowNativePasswords: true,
		ParseTime:            true,
	},
	Server: ServerConfig{
		Host: "0.0.0.0",
		Port: "80",
	},
}
