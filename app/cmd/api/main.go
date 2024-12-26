package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yaken-org/hakushi/cmd/api/handler/healthcheck"
	"github.com/yaken-org/hakushi/pkg/config"
	"github.com/yaken-org/hakushi/pkg/server"
)

func main() {
	s := server.New()
	for _, r := range routes {
		s.Add(r.method, r.path, r.handler)
	}
	c := config.Get()
	e := s.Echo()

	// Graceful shutdown
	// ref: https://echo.labstack.com/docs/cookbook/graceful-shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		host := c.Server.Host + ":" + c.Server.Port
		if err := e.Start(host); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

var routes = []struct {
	method  string
	path    string
	handler server.HandlerFunc
}{
	// Healthcheck
	{"GET", "/api/v1/health", healthcheck.Health},
}
