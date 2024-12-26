package api

import (
	"github.com/yaken-org/hakushi/pkg/api/handler"
	"github.com/yaken-org/hakushi/pkg/config"
	"github.com/yaken-org/hakushi/pkg/server"
)

func Start() {
	s := server.New()
	for _, r := range routes {
		s.Add(r.method, r.path, r.handler)
	}
	c := config.Get()
	s.Core().Logger.Fatal(s.Start(c.Server.Host + ":" + c.Server.Port))
}

var routes = []struct {
	method  string
	path    string
	handler server.HandlerFunc
}{
	{"GET", "/api/v1/health", handler.Health},
}
