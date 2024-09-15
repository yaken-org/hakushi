package handler

import (
	"log/slog"
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

	postIdToTags, err := service.FindPostRelatedTagsByPostIDs(postIDs)
	if err != nil {
		return err
	}

	var apiPosts []*model.APIPost
	for _, post := range posts {
		apiPost := post.ToAPIPost(postIdToAnnotations[post.ID], postIdToTags[post.ID])
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

	tags, err := service.FindPostRelatedTags(post.ID)
	if err != nil {
		return err
	}

	apiPost := post.ToAPIPost(annotations, tags)
	return c.JSON(200, apiPost)
}

// CreatePost はポストを作成するハンドラ
func CreatePost(c echo.Context) error {
	apiPost := new(model.APIPost)
	if err := c.Bind(apiPost); err != nil {
		slog.Error(err.Error())
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
	annotations := []*model.Annotation{}
	for _, annotation := range apiPost.Annotations {
		annotation, err = service.CreateAnnotation(post, annotation.ProductID, annotation.DisplayName, annotation.X, annotation.Y)
		if err != nil {
			continue
		}
		annotations = append(annotations, annotation)
	}

	tags, err := service.CreateTagAndPostTagMapping(post.ID, apiPost.Tags)
	if err != nil {
		return err
	}

	// API で通信する形のポストに整形して返す
	return c.JSON(200, post.ToAPIPost(annotations, tags))
}

// GetUserPosts はユーザのポストを取得するハンドラ
func GetUserPosts(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// ユーザのポストを取得する
	posts, err := service.FindPostsByUserAccountID(id)
	if err != nil {
		return err
	}

	// ポストからIDを取得する
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	// ポストからアノテーションを取得する
	postIdToAnnotations, err := service.FindAnnotationsByPostIDs(postIDs)
	if err != nil {
		return err
	}
	// ポストからタグを取得する
	postIdToTags, err := service.FindPostRelatedTagsByPostIDs(postIDs)
	if err != nil {
		return err
	}

	// API で通信する形のポストに整形して返す
	apiPosts := []*model.APIPost{}
	for _, post := range posts {
		apiPost := post.ToAPIPost(postIdToAnnotations[post.ID], postIdToTags[post.ID])
		apiPosts = append(apiPosts, apiPost)
	}

	return c.JSON(200, apiPosts)
}

func GetPostTags(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	tags, err := service.FindPostRelatedTags(postID)
	if err != nil {
		return err
	}

	return c.JSON(200, tags)
}

func CreatePostTag(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tagID, err := strconv.ParseInt(c.Param("tag_id"), 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	postTag, err := service.CreatePostTag(postID, tagID)
	if err != nil {
		return err
	}

	return c.JSON(200, postTag)
}
