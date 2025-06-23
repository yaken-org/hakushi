package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yaken-org/hakushi/internal/usecase"
)

type MockPostUsecase struct {
	mock.Mock
}

func (m *MockPostUsecase) GetPostByID(ctx context.Context, id int64) (*usecase.PostOutput, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.PostOutput), args.Error(1)
}

func (m *MockPostUsecase) GetAllPosts(ctx context.Context) ([]*usecase.PostOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*usecase.PostOutput), args.Error(1)
}

func (m *MockPostUsecase) CreatePost(ctx context.Context, input usecase.CreatePostInput) (*usecase.PostOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.PostOutput), args.Error(1)
}

func (m *MockPostUsecase) UpdatePost(ctx context.Context, input usecase.UpdatePostInput) (*usecase.PostOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.PostOutput), args.Error(1)
}

func (m *MockPostUsecase) DeletePost(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPostUsecase) LikePost(ctx context.Context, id int64) (int, error) {
	args := m.Called(ctx, id)
	return args.Int(0), args.Error(1)
}

func (m *MockPostUsecase) GetPostsByUserID(ctx context.Context, userID int64) ([]*usecase.PostOutput, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*usecase.PostOutput), args.Error(1)
}

func (m *MockPostUsecase) GetPostsOrderByLikes(ctx context.Context) ([]*usecase.PostOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*usecase.PostOutput), args.Error(1)
}

func TestPostHandler_GetPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/post/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		expectedPost := &usecase.PostOutput{
			ID:      1,
			Title:   "Test Post",
			Content: "Test Content",
		}

		mockUsecase.On("GetPostByID", mock.Anything, int64(1)).Return(expectedPost, nil)

		err := handler.GetPost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response usecase.PostOutput
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedPost.ID, response.ID)
		assert.Equal(t, expectedPost.Title, response.Title)

		mockUsecase.AssertExpectations(t)
	})

	t.Run("invalid id", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/post/invalid", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid")

		err := handler.GetPost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response ErrorResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid post ID", response.Error)
	})

	t.Run("post not found", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/post/999", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("999")

		mockUsecase.On("GetPostByID", mock.Anything, int64(999)).Return(nil, errors.New("post not found"))

		err := handler.GetPost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		mockUsecase.AssertExpectations(t)
	})
}

func TestPostHandler_CreatePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		requestBody := CreatePostRequest{
			UserAccountID: "1",
			ImageID:       100,
			Title:         "Test Post",
			Content:       "Test Content",
			Annotations:   []CreateAnnotationRequest{},
			Tags:          []string{"test"},
		}

		body, _ := json.Marshal(requestBody)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedPost := &usecase.PostOutput{
			ID:      1,
			Title:   "Test Post",
			Content: "Test Content",
		}

		mockUsecase.On("CreatePost", mock.Anything, mock.AnythingOfType("usecase.CreatePostInput")).Return(expectedPost, nil)

		err := handler.CreatePost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response usecase.PostOutput
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedPost.ID, response.ID)

		mockUsecase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader([]byte("invalid json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.CreatePost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		requestBody := CreatePostRequest{
			UserAccountID: "",
			ImageID:       0,
			Title:         "",
			Content:       "",
		}

		body, _ := json.Marshal(requestBody)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.CreatePost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response ErrorResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Validation failed", response.Error)
	})
}

func TestPostHandler_SendLikeToPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := &MockPostUsecase{}
		handler := NewPostHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/post/1/like", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockUsecase.On("LikePost", mock.Anything, int64(1)).Return(5, nil)

		err := handler.SendLikeToPost(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, float64(5), response["likes"])

		mockUsecase.AssertExpectations(t)
	})
}