package usecase

import (
	"context"
	"fmt"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type UserAccountUsecase struct {
	userRepo repository.UserAccountRepository
}

func NewUserAccountUsecase(userRepo repository.UserAccountRepository) *UserAccountUsecase {
	return &UserAccountUsecase{
		userRepo: userRepo,
	}
}

func (u *UserAccountUsecase) CreateUserAccount(ctx context.Context, input CreateUserAccountInput) (*UserAccountOutput, error) {
	userAccount := entity.NewUserAccount(
		input.Sub,
		input.Username,
		input.Email,
		input.Name,
		input.ImageURL,
	)
	
	if err := u.userRepo.Create(ctx, userAccount); err != nil {
		return nil, fmt.Errorf("failed to create user account: %w", err)
	}
	
	return &UserAccountOutput{
		ID:        userAccount.ID,
		Sub:       userAccount.Sub,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		Name:      userAccount.Name,
		ImageURL:  userAccount.ImageURL,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
	}, nil
}

func (u *UserAccountUsecase) GetUserAccountByID(ctx context.Context, id int64) (*UserAccountOutput, error) {
	userAccount, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user account: %w", err)
	}
	
	return &UserAccountOutput{
		ID:        userAccount.ID,
		Sub:       userAccount.Sub,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		Name:      userAccount.Name,
		ImageURL:  userAccount.ImageURL,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
	}, nil
}

func (u *UserAccountUsecase) GetUserAccountBySub(ctx context.Context, sub string) (*UserAccountOutput, error) {
	userAccount, err := u.userRepo.FindBySub(ctx, sub)
	if err != nil {
		return nil, fmt.Errorf("failed to get user account: %w", err)
	}
	
	return &UserAccountOutput{
		ID:        userAccount.ID,
		Sub:       userAccount.Sub,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		Name:      userAccount.Name,
		ImageURL:  userAccount.ImageURL,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
	}, nil
}

func (u *UserAccountUsecase) GetUserAccountByUsername(ctx context.Context, username string) (*UserAccountOutput, error) {
	userAccount, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user account: %w", err)
	}
	
	return &UserAccountOutput{
		ID:        userAccount.ID,
		Sub:       userAccount.Sub,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		Name:      userAccount.Name,
		ImageURL:  userAccount.ImageURL,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
	}, nil
}

func (u *UserAccountUsecase) UpdateUserAccount(ctx context.Context, input UpdateUserAccountInput) (*UserAccountOutput, error) {
	userAccount, err := u.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user account: %w", err)
	}
	
	userAccount.Username = input.Username
	userAccount.Email = input.Email
	userAccount.Name = input.Name
	userAccount.ImageURL = input.ImageURL
	
	if err := u.userRepo.Update(ctx, userAccount); err != nil {
		return nil, fmt.Errorf("failed to update user account: %w", err)
	}
	
	return &UserAccountOutput{
		ID:        userAccount.ID,
		Sub:       userAccount.Sub,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		Name:      userAccount.Name,
		ImageURL:  userAccount.ImageURL,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
	}, nil
}