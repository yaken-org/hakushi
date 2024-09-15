package server

import "github.com/yaken-org/hakushi/internal/server/handler"

func (s *Server) configureRoute() {
	e := s.Engine

	e.GET("/health", handler.Health)
}
