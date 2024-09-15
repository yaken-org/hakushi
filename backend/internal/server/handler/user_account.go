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
		return c.NoContent(http.StatusBadRequest)
	}

	account, err := service.CreateUserAccount(
		user.Name,
		user.DisplayName,
		user.IconURL,
		user.Sub,
	)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, account)
}

func GetUserAcocunt(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	account, err := service.FindUserAccountById(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, account)
}
