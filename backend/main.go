package main

import (
	"log/slog"
	"os"

	"github.com/yaken-org/hakushi/internal/config"
	"github.com/yaken-org/hakushi/internal/infrastructure/database"
	"github.com/yaken-org/hakushi/internal/wire"
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
	cfg := config.New(e)

	slog.Debug("Initializing database")
	db, err := database.New(cfg)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Debug("Initializing server with dependency injection")
	server, err := wire.InitializeServer(cfg, db)
	if err != nil {
		slog.Error("Failed to initialize server", "error", err)
		os.Exit(1)
	}

	slog.Info("Starting server")
	if err := server.Start(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
