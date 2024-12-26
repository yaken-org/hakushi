package config

var test = Config{
	Database: DatabaseConfig{
		User:                 "root",
		Password:             "root",
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
