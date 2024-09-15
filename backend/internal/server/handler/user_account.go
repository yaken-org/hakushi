package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/model"
	"github.com/yaken-org/hakushi/internal/service"
)

func CreateUserAccount(c echo.Context) error {
	user := new(model.UserAccount)
	if err := c.Bind(user); err != nil {
		return err
	}

	account, err := service.CreateUserAccount(
		user.Name,
		user.DisplayName,
		user.IconURL,
		user.Sub,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}

func GetUserAcocunt(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	account, err := service.FindUserAccountById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}

func GetUserAccountBySub(c echo.Context) error {
	sub := c.Param("id")

	account, err := service.FindUserAccountBySub(sub)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}
