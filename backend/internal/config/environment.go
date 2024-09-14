package config

import (
	"io"
	"os"
	"path/filepath"
)

type Environment interface {
	Name() string
	filepath() string
	reader() io.Reader
}

// path は環境に応じたファイルパスを返す。
func path(e Environment) string {
	return filepath.Join("config", e.Name()+".yaml")
}

// file は環境に応じたファイルを開いて io.Reader を返す。
func file(e Environment) io.Reader {
	f, err := os.Open(e.filepath())
	if err != nil {
		panic(err)
	}
	return f
}

// Production は本番環境を表す Environment を返す。
func Production() Environment {
	return new(EnvironmentProduction)
}

type EnvironmentProduction struct{}

func (e *EnvironmentProduction) Name() string {
	return "production"
}

func (e *EnvironmentProduction) filepath() string {
	return path(e)
}

func (e *EnvironmentProduction) reader() io.Reader {
	return file(e)
}

// Development は開発環境を表す Environment を返す。
func Development() Environment {
	return new(EnvironmentDevelopment)
}

type EnvironmentDevelopment struct{}

func (e *EnvironmentDevelopment) Name() string {
	return "development"
}

func (e *EnvironmentDevelopment) filepath() string {
	return path(e)
}

func (e *EnvironmentDevelopment) reader() io.Reader {
	return file(e)
}
