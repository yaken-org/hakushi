package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/service"
)

func GetAllTags(c echo.Context) error {
	tags, err := service.FindAllTags()
	if err != nil {
		return err
	}
	return c.JSON(200, tags)
}

func GetTaggedPosts(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tag, err := service.FindTagByID(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	posts, err := service.FindPostsByTag(tag)
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(200, posts)
}
