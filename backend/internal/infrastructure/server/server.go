package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yaken-org/hakushi/internal/adapter/handler"
	"github.com/yaken-org/hakushi/internal/config"
)

type Server struct {
	engine      *echo.Echo
	config      *config.Config
	postHandler *handler.PostHandler
	userHandler *handler.UserAccountHandler
}

func New(
	config *config.Config,
	postHandler *handler.PostHandler,
	userHandler *handler.UserAccountHandler,
) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := &Server{
		engine:      e,
		config:      config,
		postHandler: postHandler,
		userHandler: userHandler,
	}

	s.configureRoutes()
	return s
}

func (s *Server) configureRoutes() {
	e := s.engine

	e.GET("/health", handler.Health)

	api := e.Group("/api")

	// User account routes
	api.POST("/account", s.userHandler.CreateUserAccount)
	api.GET("/account/:id", s.userHandler.GetUserAccount)
	api.PUT("/account/:id", s.userHandler.UpdateUserAccount)
	api.GET("/account/:id/posts", s.postHandler.GetUserPosts)
	api.GET("/account/sub/:id", s.userHandler.GetUserAccountBySub)
	api.GET("/account/name/:name", s.userHandler.GetUserAccountByName)

	// Post routes
	api.GET("/post", s.postHandler.GetAllPosts)
	api.POST("/post", s.postHandler.CreatePost)
	api.GET("/post/:id", s.postHandler.GetPost)
	api.PUT("/post/:id", s.postHandler.UpdatePost)
	api.DELETE("/post/:id", s.postHandler.DeletePost)
	api.POST("/post/:id/like", s.postHandler.SendLikeToPost)
}

func (s *Server) Start() error {
	host := s.config.Server.Host
	port := s.config.Server.Port
	address := fmt.Sprintf("%s:%d", host, port)
	return s.engine.Start(address)
}
