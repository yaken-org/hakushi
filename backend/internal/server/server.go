package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yaken-org/hakushi/internal/config"
)

var server *Server

type Server struct {
	Engine *echo.Echo
	Config *config.Config
}

func Get() *Server {
	if server == nil {
		panic("server is not initialized")
	}

	return server
}

func Initialize(config *config.Config) error {
	if server != nil {
		return nil
	}

	s := new(Server)
	s.Config = config
	s.Engine = echo.New()

	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())

	s.configureRoute()

	server = s
	return nil
}

func (s *Server) Start() error {
	host := s.Config.Server.Host
	port := s.Config.Server.Port
	address := fmt.Sprintf("%s:%d", host, port)
	return s.Engine.Start(address)
}
