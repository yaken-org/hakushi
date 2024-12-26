package server

import (
	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/pkg/database"
)

type Context struct {
	echo.Context
	DB *database.Database
}

func NewContext(c echo.Context) Context {
	return Context{c, database.New()}
}
