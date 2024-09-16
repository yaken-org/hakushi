package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/service"
)

func Search(c echo.Context) error {
	searchWord := c.QueryParam("word")
	if searchWord == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "search word is missing.",
		})
	}
	slog.Info(fmt.Sprintf("search word: %s", searchWord))

	posts, err := service.FindPostByNameRough(searchWord)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, posts)
}
