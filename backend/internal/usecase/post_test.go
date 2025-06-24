package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/testutil/fixture"
	mockRepo "github.com/yaken-org/hakushi/internal/testutil/mock"
)

func TestPostUsecase_GetPostByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		postID := int64(1)
		expectedPost := fixture.NewPost(1, 1, 100, "Test Post", "Test Content")
		expectedAnnotations := []*entity.Annotation{
			fixture.NewAnnotation(1, 1, nil, "Test Annotation", 100, 200),
		}
		expectedTags := []*entity.Tag{
			fixture.NewTag(1, "test"),
		}

		mockPostRepo.On("FindByID", ctx, postID).Return(expectedPost, nil)
		mockAnnotationRepo.On("FindByPostID", ctx, postID).Return(expectedAnnotations, nil)
		mockTagRepo.On("FindByPostID", ctx, postID).Return(expectedTags, nil)

		result, err := usecase.GetPostByID(ctx, postID)

		assert.NoError(t, err)
		assert.Equal(t, expectedPost.ID, result.ID)
		assert.Equal(t, expectedPost.Title, result.Title)
		assert.Equal(t, expectedPost.Content, result.Content)
		assert.Len(t, result.Annotations, 1)
		assert.Len(t, result.Tags, 1)

		mockPostRepo.AssertExpectations(t)
		mockAnnotationRepo.AssertExpectations(t)
		mockTagRepo.AssertExpectations(t)
	})

	t.Run("post not found", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		postID := int64(999)

		mockPostRepo.On("FindByID", ctx, postID).Return(nil, errors.New("post not found"))

		result, err := usecase.GetPostByID(ctx, postID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to get post")

		mockPostRepo.AssertExpectations(t)
	})
}

func TestPostUsecase_CreatePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		input := CreatePostInput{
			UserAccountID: 1,
			ImageID:       100,
			Title:         "Test Post",
			Content:       "Test Content",
			Annotations: []CreateAnnotationInput{
				{ProductID: nil, DisplayName: "Test Annotation", X: 100, Y: 200},
			},
			Tags: []string{"test", "sample"},
		}
		expectedUser := fixture.NewUserAccount(1, "sub123", "testuser", "test@example.com", "Test User", "")

		mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(context.Context) error")).Return(nil)
		mockUserRepo.On("FindByID", ctx, input.UserAccountID).Return(expectedUser, nil)
		mockPostRepo.On("Create", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Run(func(args mock.Arguments) {
			post := args.Get(1).(*entity.Post)
			post.ID = 123
		})
		mockAnnotationRepo.On("Create", ctx, mock.AnythingOfType("*entity.Annotation")).Return(nil).Run(func(args mock.Arguments) {
			annotation := args.Get(1).(*entity.Annotation)
			annotation.ID = 456
		})
		mockTagRepo.On("FindByName", ctx, "test").Return(nil, errors.New("not found"))
		mockTagRepo.On("Create", ctx, mock.AnythingOfType("*entity.Tag")).Return(nil).Run(func(args mock.Arguments) {
			tag := args.Get(1).(*entity.Tag)
			tag.ID = 789
		})
		mockTagRepo.On("FindByName", ctx, "sample").Return(nil, errors.New("not found"))
		mockTagRepo.On("Create", ctx, mock.AnythingOfType("*entity.Tag")).Return(nil).Run(func(args mock.Arguments) {
			tag := args.Get(1).(*entity.Tag)
			tag.ID = 790
		})
		mockPostTagRepo.On("Create", ctx, mock.AnythingOfType("*entity.PostTag")).Return(nil).Twice()

		result, err := usecase.CreatePost(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Title, result.Title)
		assert.Equal(t, input.Content, result.Content)

		mockTxManager.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		input := CreatePostInput{
			UserAccountID: 999,
			ImageID:       100,
			Title:         "Test Post",
			Content:       "Test Content",
			Annotations:   []CreateAnnotationInput{},
			Tags:          []string{},
		}

		mockTxManager.On("WithTransaction", ctx, mock.AnythingOfType("func(context.Context) error")).Return(errors.New("user not found"))
		mockUserRepo.On("FindByID", ctx, input.UserAccountID).Return(nil, errors.New("user not found"))

		result, err := usecase.CreatePost(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockTxManager.AssertExpectations(t)
	})
}

func TestPostUsecase_LikePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		postID := int64(1)
		post := fixture.NewPost(1, 1, 100, "Test Post", "Test Content")
		originalLikes := post.Likes

		mockPostRepo.On("FindByID", ctx, postID).Return(post, nil)
		mockPostRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil)

		likes, err := usecase.LikePost(ctx, postID)

		assert.NoError(t, err)
		assert.Equal(t, originalLikes+1, likes)

		mockPostRepo.AssertExpectations(t)
	})
}

func TestPostUsecase_GetAllPosts(t *testing.T) {
	t.Run("success with multiple posts", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()
		expectedPosts := []*entity.Post{
			fixture.NewPost(1, 1, 100, "Post 1", "Content 1"),
			fixture.NewPost(2, 1, 101, "Post 2", "Content 2"),
		}

		mockPostRepo.On("FindAll", ctx).Return(expectedPosts, nil)
		mockAnnotationRepo.On("FindByPostIDs", ctx, []int64{1, 2}).Return(map[int64][]*entity.Annotation{}, nil)
		mockTagRepo.On("FindByPostIDs", ctx, []int64{1, 2}).Return(map[int64][]*entity.Tag{}, nil)

		result, err := usecase.GetAllPosts(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Post 1", result[0].Title)
		assert.Equal(t, "Post 2", result[1].Title)

		mockPostRepo.AssertExpectations(t)
		mockAnnotationRepo.AssertExpectations(t)
		mockTagRepo.AssertExpectations(t)
	})

	t.Run("empty result", func(t *testing.T) {
		mockPostRepo := &mockRepo.PostRepository{}
		mockUserRepo := &mockRepo.UserAccountRepository{}
		mockAnnotationRepo := &mockRepo.AnnotationRepository{}
		mockTagRepo := &mockRepo.TagRepository{}
		mockPostTagRepo := &mockRepo.PostTagRepository{}
		mockTxManager := &mockRepo.TransactionManager{}

		usecase := NewPostUsecase(
			mockPostRepo,
			mockUserRepo,
			mockAnnotationRepo,
			mockTagRepo,
			mockPostTagRepo,
			mockTxManager,
		)

		ctx := context.Background()

		mockPostRepo.On("FindAll", ctx).Return([]*entity.Post{}, nil)

		result, err := usecase.GetAllPosts(ctx)

		assert.NoError(t, err)
		assert.Empty(t, result)

		mockPostRepo.AssertExpectations(t)
	})
}