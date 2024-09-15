package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/internal/model"
	"github.com/yaken-org/hakushi/internal/service"
)

func GetAllPosts(c echo.Context) error {
	posts, err := service.FindAllPosts()
	if err != nil {
		return err
	}
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	postIdToAnnotations, err := service.FindAnnotationsByPostIDs(postIDs)
	if err != nil {
		return err
	}

	var apiPosts []*model.APIPost
	for _, post := range posts {
		apiPost := post.ToAPIPost(postIdToAnnotations[post.ID])
		apiPosts = append(apiPosts, apiPost)
	}

	return c.JSON(200, apiPosts)
}

func GetPost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	post, err := service.FindPostByID(id)
	if err != nil {
		return err
	}

	annotations := make([]*model.Annotation, 0)
	annotations, err = service.FindAnnotationsByPostID(post.ID)
	if err != nil {
		return err
	}
	apiPost := post.ToAPIPost(annotations)
	return c.JSON(200, apiPost)
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
