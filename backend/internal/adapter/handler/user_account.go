package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/usecase"
)

type UserAccountHandler struct {
	userUsecase usecase.UserAccountUsecaseInterface
}

func NewUserAccountHandler(userUsecase usecase.UserAccountUsecaseInterface) *UserAccountHandler {
	return &UserAccountHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserAccountHandler) CreateUserAccount(c echo.Context) error {
	var req CreateUserAccountRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	input := req.ToUsecaseInput()
	userAccount, err := h.userUsecase.CreateUserAccount(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to create user account",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, userAccount)
}

func (h *UserAccountHandler) GetUserAccount(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a number",
		})
	}

	userAccount, err := h.userUsecase.GetUserAccountByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User account not found",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userAccount)
}

func (h *UserAccountHandler) GetUserAccountBySub(c echo.Context) error {
	sub := c.Param("id")
	if sub == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid sub",
			Message: "Sub parameter is required",
		})
	}

	userAccount, err := h.userUsecase.GetUserAccountBySub(c.Request().Context(), sub)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User account not found",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userAccount)
}

func (h *UserAccountHandler) GetUserAccountByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid name",
			Message: "Name parameter is required",
		})
	}

	userAccount, err := h.userUsecase.GetUserAccountByUsername(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User account not found",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userAccount)
}

func (h *UserAccountHandler) UpdateUserAccount(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a number",
		})
	}

	var req UpdateUserAccountRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	input := usecase.UpdateUserAccountInput{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
		ImageURL: req.ImageURL,
	}

	userAccount, err := h.userUsecase.UpdateUserAccount(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to update user account",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userAccount)
}