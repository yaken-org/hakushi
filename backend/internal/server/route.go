package server

import "github.com/yaken-org/hakushi/internal/server/handler"

func (s *Server) configureRoute() {
	e := s.Engine

	e.GET("/health", handler.Health)

	api := e.Group("/api")
	api.POST("/account", handler.CreateUserAccount)
	api.GET("/account/:id", handler.GetUserAcocunt)
	api.GET("/account/:id/posts", handler.GetUserPosts)

	api.GET("/post", handler.GetAllPosts)
	api.POST("/post", handler.CreatePost)
	api.GET("/post/:id", handler.GetPost)
	api.GET("/post/:id/tags", handler.GetPostTags)
}
