package config

type Schema struct {
	Server   SchemaServer   `yaml:"server"`
	Database SchemaDatabase `yaml:"database"`
}

type SchemaServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SchemaDatabase struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	Timezone        string `yaml:"timezone"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

func DefaultSchema() *Schema {
	return &Schema{
		Server: SchemaServer{
			Host: "",
			Port: 0,
		},
		Database: SchemaDatabase{
			Host:            "",
			Port:            0,
			Username:        "",
			Password:        "",
			Database:        "",
			Timezone:        "Asia/Tokyo",
			MaxIdleConns:    10,
			MaxOpenConns:    10,
			ConnMaxLifetime: 10,
		},
	}
}
