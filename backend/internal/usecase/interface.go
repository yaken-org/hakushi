package usecase

import "context"

type PostUsecaseInterface interface {
	GetPostByID(ctx context.Context, id int64) (*PostOutput, error)
	GetAllPosts(ctx context.Context) ([]*PostOutput, error)
	CreatePost(ctx context.Context, input CreatePostInput) (*PostOutput, error)
	UpdatePost(ctx context.Context, input UpdatePostInput) (*PostOutput, error)
	DeletePost(ctx context.Context, id int64) error
	LikePost(ctx context.Context, id int64) (int, error)
	GetPostsByUserID(ctx context.Context, userID int64) ([]*PostOutput, error)
	GetPostsOrderByLikes(ctx context.Context) ([]*PostOutput, error)
}

type UserAccountUsecaseInterface interface {
	CreateUserAccount(ctx context.Context, input CreateUserAccountInput) (*UserAccountOutput, error)
	GetUserAccountByID(ctx context.Context, id int64) (*UserAccountOutput, error)
	GetUserAccountBySub(ctx context.Context, sub string) (*UserAccountOutput, error)
	GetUserAccountByUsername(ctx context.Context, username string) (*UserAccountOutput, error)
	UpdateUserAccount(ctx context.Context, input UpdateUserAccountInput) (*UserAccountOutput, error)
}