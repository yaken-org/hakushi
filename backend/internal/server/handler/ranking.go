package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/service"
)

func GetRanking(c echo.Context) error {
	category := c.QueryParam("category")
	if category == "" {
		category = "like"
	}

	posts, err := service.FindPostsOrderByLikes()
	if err != nil {
		return err
	}

	return c.JSON(200, posts)
}
