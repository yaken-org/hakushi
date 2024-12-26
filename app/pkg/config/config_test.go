package config

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		want        Config
	}{
		{
			name:        "APP_ENVが local の時は local 用の設定が返る",
			environment: "local",
			want:        local,
		},
		{
			name:        "APP_ENVが development の時は development 用の設定が返る",
			environment: "development",
			want:        development,
		},
		{
			name:        "APP_ENVが production の時は production 用の設定が返る",
			environment: "production",
			want:        production,
		},
		{
			name:        "APP_ENVが空文字の時は local 用の設定が返る",
			environment: "",
			want:        local,
		},
		{
			name:        "APP_ENVが存在しない時は何も定義されていない設定が返る",
			environment: "unknown",
			want:        Config{},
		},
	}

	for _, tt := range tests {
		err := os.Unsetenv("APP_ENV")
		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv("APP_ENV", tt.environment)
			if err != nil {
				t.Fatal(err)
			}
			cfg = Config{}

			if got := Get(); got != tt.want {
				t.Errorf("\nGet() = %#v\nwant = %#v", got, tt.want)
			}
		})
	}
}

func TestSetEnv(t *testing.T) {
	newEnvConfig := Config{
		Database: DatabaseConfig{
			User: "new_user",
		},
	}

	localEnvConfig := Config{
		Database: DatabaseConfig{
			User: "local_user",
		},
	}

	tests := []struct {
		name        string
		environment string
		cfg         Config
		want        Config
	}{
		{
			name:        "新たな環境別の設定を追加できる",
			environment: "new",
			cfg:         newEnvConfig,
			want:        newEnvConfig,
		},
		{
			name:        "既存の環境別の設定を上書きできる",
			environment: "local",
			cfg:         localEnvConfig,
			want:        localEnvConfig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetEnv(tt.environment, tt.cfg)
			if got := env[tt.environment]; got != tt.want {
				t.Errorf("SetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
