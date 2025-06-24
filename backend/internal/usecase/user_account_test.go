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

func TestUserAccountUsecase_CreateUserAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		input := CreateUserAccountInput{
			Sub:      "google-oauth2|123456",
			Username: "testuser",
			Email:    "test@example.com",
			Name:     "Test User",
			ImageURL: "https://example.com/avatar.jpg",
		}

		mockUserRepo.On("Create", ctx, mock.AnythingOfType("*entity.UserAccount")).Return(nil).Run(func(args mock.Arguments) {
			userAccount := args.Get(1).(*entity.UserAccount)
			userAccount.ID = 123
		})

		result, err := usecase.CreateUserAccount(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int64(123), result.ID)
		assert.Equal(t, input.Sub, result.Sub)
		assert.Equal(t, input.Username, result.Username)
		assert.Equal(t, input.Email, result.Email)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.ImageURL, result.ImageURL)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		input := CreateUserAccountInput{
			Sub:      "google-oauth2|123456",
			Username: "testuser",
			Email:    "test@example.com",
			Name:     "Test User",
			ImageURL: "",
		}

		mockUserRepo.On("Create", ctx, mock.AnythingOfType("*entity.UserAccount")).Return(errors.New("database error"))

		result, err := usecase.CreateUserAccount(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to create user account")

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserAccountUsecase_GetUserAccountByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		userID := int64(1)
		expectedUser := fixture.NewUserAccount(
			1,
			"google-oauth2|123456",
			"testuser",
			"test@example.com",
			"Test User",
			"https://example.com/avatar.jpg",
		)

		mockUserRepo.On("FindByID", ctx, userID).Return(expectedUser, nil)

		result, err := usecase.GetUserAccountByID(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, result.ID)
		assert.Equal(t, expectedUser.Sub, result.Sub)
		assert.Equal(t, expectedUser.Username, result.Username)
		assert.Equal(t, expectedUser.Email, result.Email)
		assert.Equal(t, expectedUser.Name, result.Name)
		assert.Equal(t, expectedUser.ImageURL, result.ImageURL)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		userID := int64(999)

		mockUserRepo.On("FindByID", ctx, userID).Return(nil, errors.New("user not found"))

		result, err := usecase.GetUserAccountByID(ctx, userID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to get user account")

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserAccountUsecase_GetUserAccountBySub(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		sub := "google-oauth2|123456"
		expectedUser := fixture.NewUserAccount(
			1,
			sub,
			"testuser",
			"test@example.com",
			"Test User",
			"",
		)

		mockUserRepo.On("FindBySub", ctx, sub).Return(expectedUser, nil)

		result, err := usecase.GetUserAccountBySub(ctx, sub)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Sub, result.Sub)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserAccountUsecase_UpdateUserAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo := &mockRepo.UserAccountRepository{}
		usecase := NewUserAccountUsecase(mockUserRepo)

		ctx := context.Background()
		userID := int64(1)
		input := UpdateUserAccountInput{
			ID:       userID,
			Username: "updateduser",
			Email:    "updated@example.com",
			Name:     "Updated User",
			ImageURL: "https://example.com/new-avatar.jpg",
		}

		existingUser := fixture.NewUserAccount(
			userID,
			"google-oauth2|123456",
			"olduser",
			"old@example.com",
			"Old User",
			"",
		)

		mockUserRepo.On("FindByID", ctx, userID).Return(existingUser, nil)
		mockUserRepo.On("Update", ctx, mock.AnythingOfType("*entity.UserAccount")).Return(nil)

		result, err := usecase.UpdateUserAccount(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, input.Username, result.Username)
		assert.Equal(t, input.Email, result.Email)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.ImageURL, result.ImageURL)

		mockUserRepo.AssertExpectations(t)
	})
}