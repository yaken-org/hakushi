package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/model"
	"github.com/yaken-org/hakushi/internal/service"
)

func GetAllPosts(c echo.Context) error {
	posts, err := service.FindAllPosts()
	if err != nil {
		return err
	}
	return c.JSON(200, posts)
}

func GetPost(c echo.Context) error {
	return nil
}

func CreatePost(c echo.Context) error {
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return err
	}

	// userAccount := c.Get("user_account").(*model.UserAccount)
	userAccount := new(model.UserAccount)
	userAccount.ID = 1

	post, err := service.CreatePost(*userAccount, post.Title, post.Content)
	if err != nil {
		return err
	}
	return c.JSON(200, post)
}
