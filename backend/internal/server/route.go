package server

import "github.com/yaken-org/hakushi/internal/server/routes"

func (s *Server) configureRoute() {
	e := s.Engine

	e.GET("/health", routes.Health)
}
