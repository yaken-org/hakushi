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
	apiPost := new(model.APIPost)
	if err := c.Bind(apiPost); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// アカウント情報を取得する
	userAccountID := apiPost.UserAccountID
	account, err := service.FindUserAccountById(userAccountID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	// ポストを作成する
	post, err := service.CreatePost(*account, apiPost.ImageID, apiPost.Title, apiPost.Content)
	if err != nil {
		return err
	}

	// ポストに紐づけられた仮のアノテーションを登録する
	annotations := apiPost.Annotations
	if annotations == nil {
		annotations = []*model.Annotation{}
	}
	for _, annotation := range annotations {
		annotation, err = service.CreateAnnotation(post, annotation.ProductID, annotation.DisplayName, annotation.X, annotation.Y)
		if err != nil {
			continue
		}
		annotations = append(annotations, annotation)
	}

	// API で通信する形のポストに整形して返す
	return c.JSON(200, post.ToAPIPost(annotations))
}
