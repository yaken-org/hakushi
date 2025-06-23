package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/usecase"
)

type PostHandler struct {
	postUsecase *usecase.PostUsecase
}

func NewPostHandler(postUsecase *usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
	}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	posts, err := h.postUsecase.GetAllPosts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to get posts",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid post ID",
			Message: "Post ID must be a number",
		})
	}

	post, err := h.postUsecase.GetPostByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "Post not found",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	var req CreatePostRequest
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
	post, err := h.postUsecase.CreatePost(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to create post",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetUserPosts(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a number",
		})
	}

	posts, err := h.postUsecase.GetPostsByUserID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to get user posts",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) SendLikeToPost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid post ID",
			Message: "Post ID must be a number",
		})
	}

	likes, err := h.postUsecase.LikePost(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to like post",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"likes": likes,
	})
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid post ID",
			Message: "Post ID must be a number",
		})
	}

	var req UpdatePostRequest
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

	input := usecase.UpdatePostInput{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}

	post, err := h.postUsecase.UpdatePost(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to update post",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid post ID",
			Message: "Post ID must be a number",
		})
	}

	if err := h.postUsecase.DeletePost(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to delete post",
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}