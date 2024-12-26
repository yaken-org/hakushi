package database

import (
	"testing"

	"github.com/yaken-org/hakushi/pkg/config"
)

func Test_dsn(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常系",
			args: args{
				cfg: config.Config{
					Database: config.DatabaseConfig{
						User:                 "user",
						Password:             "password",
						Host:                 "host",
						Port:                 "port",
						Database:             "database",
						TimeZone:             "Asia/Tokyo",
						AllowNativePasswords: true,
						ParseTime:            true,
					},
				},
			},
			want: "user:password@tcp(host:port)/database?checkConnLiveness=false&loc=Asia%2FTokyo&parseTime=true&maxAllowedPacket=0",
		},
		{
			name: "タイムゾーンが不正な場合はUTCになる",
			args: args{
				cfg: config.Config{
					Database: config.DatabaseConfig{
						User:                 "user",
						Password:             "password",
						Host:                 "host",
						Port:                 "port",
						Database:             "database",
						TimeZone:             "invalid",
						AllowNativePasswords: true,
						ParseTime:            true,
					},
				},
			},
			// UTC の場合は、DSN に loc が含まれない
			want: "user:password@tcp(host:port)/database?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dsn(tt.args.cfg); got != tt.want {
				t.Errorf("dsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
