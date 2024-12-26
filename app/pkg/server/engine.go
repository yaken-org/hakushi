package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerFunc func(c Context) error

func wrap(f HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := newContext(c)
		return f(ctx)
	}
}

type Engine interface {
	Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Start(address string) error
	Core() *echo.Echo
}

type engine struct {
	*echo.Echo
}

func New() Engine {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &engine{e}
}

func (e *engine) Core() *echo.Echo {
	return e.Echo
}

func (e *engine) Start(address string) error {
	return e.Echo.Start(address)
}

func (e *engine) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.Add(method, path, wrap(handler), middleware...)
}
