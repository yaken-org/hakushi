package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerFunc func(c Context) error

func wrap(f HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := NewContext(c)
		return f(ctx)
	}
}

type Engine interface {
	Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Start(address string) error
	Echo() *echo.Echo
}

type engine struct {
	echo *echo.Echo
}

func New() Engine {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &engine{echo: e}
}

func (e *engine) Echo() *echo.Echo {
	return e.echo
}

func (e *engine) Start(address string) error {
	return e.echo.Start(address)
}

func (e *engine) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return e.echo.Add(method, path, wrap(handler), middleware...)
}
