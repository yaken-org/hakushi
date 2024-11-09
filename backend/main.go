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
	c := config.New(e)

	slog.Debug("Initializing database")
	if err := database.Initialize(c); err != nil {
		panic(err)
	}

	slog.Debug("Initializing server")
	if err := server.Initialize(c); err != nil {
		panic(err)
	}

	slog.Info("Starting server")
	if err := server.Get().Start(); err != nil {
		panic(err)
	}
}
