package main

import (
	"log/slog"
	"os"

	"github.com/yaken-org/hakushi/internal/config"
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/server"
)

func main() {
	slog.Info("Starting Hakushi")
	var e config.Environment
	if os.Getenv("APP_ENV") == "production" {
		slog.Info("Environment: Production")
		e = config.Production()
	} else {
		slog.Info("Environment: Development")
		e = config.Development()
	}

	slog.Debug("Loading configuration")
	config := config.New(e)

	slog.Debug("Initializing database")
	database.Initialize(config)

	slog.Debug("Initializing server")
	server.Initialize(config)

	slog.Info("Starting server")
	server.Get().Start()
}
