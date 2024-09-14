package config

import (
	"bufio"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	*Schema
	Environment Environment
}

// config は一つの環境に一度しか読み込まないため、
// シングルトンパターンを採用する。
var config *Config

// New は 環境に応じてコンフィグを読み込む関数。
func New(e Environment) *Config {
	if config == nil {
		config = load(e)
	}

	return config
}

// load は環境に応じたコンフィグを読み込む関数。
func load(e Environment) *Config {
	schema := DefaultSchema()
	data, err := io.ReadAll(bufio.NewReader(e.reader()))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, schema); err != nil {
		panic(err)
	}
	return &Config{
		Schema:      schema,
		Environment: e,
	}
}
