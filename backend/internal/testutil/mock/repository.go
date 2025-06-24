package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/yaken-org/hakushi/internal/domain/entity"
)

type PostRepository struct {
	mock.Mock
}

func (m *PostRepository) FindByID(ctx context.Context, id int64) (*entity.Post, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *PostRepository) FindAll(ctx context.Context) ([]*entity.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *PostRepository) FindByUserAccountID(ctx context.Context, userAccountID int64) ([]*entity.Post, error) {
	args := m.Called(ctx, userAccountID)
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *PostRepository) FindByTitleContaining(ctx context.Context, title string) ([]*entity.Post, error) {
	args := m.Called(ctx, title)
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *PostRepository) FindOrderByLikes(ctx context.Context) ([]*entity.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *PostRepository) Create(ctx context.Context, post *entity.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostRepository) Update(ctx context.Context, post *entity.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type UserAccountRepository struct {
	mock.Mock
}

func (m *UserAccountRepository) FindByID(ctx context.Context, id int64) (*entity.UserAccount, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserAccount), args.Error(1)
}

func (m *UserAccountRepository) FindBySub(ctx context.Context, sub string) (*entity.UserAccount, error) {
	args := m.Called(ctx, sub)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserAccount), args.Error(1)
}

func (m *UserAccountRepository) FindByUsername(ctx context.Context, username string) (*entity.UserAccount, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserAccount), args.Error(1)
}

func (m *UserAccountRepository) Create(ctx context.Context, userAccount *entity.UserAccount) error {
	args := m.Called(ctx, userAccount)
	return args.Error(0)
}

func (m *UserAccountRepository) Update(ctx context.Context, userAccount *entity.UserAccount) error {
	args := m.Called(ctx, userAccount)
	return args.Error(0)
}

func (m *UserAccountRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type AnnotationRepository struct {
	mock.Mock
}

func (m *AnnotationRepository) FindByID(ctx context.Context, id int64) (*entity.Annotation, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Annotation), args.Error(1)
}

func (m *AnnotationRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.Annotation, error) {
	args := m.Called(ctx, postID)
	return args.Get(0).([]*entity.Annotation), args.Error(1)
}

func (m *AnnotationRepository) FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Annotation, error) {
	args := m.Called(ctx, postIDs)
	return args.Get(0).(map[int64][]*entity.Annotation), args.Error(1)
}

func (m *AnnotationRepository) Create(ctx context.Context, annotation *entity.Annotation) error {
	args := m.Called(ctx, annotation)
	return args.Error(0)
}

func (m *AnnotationRepository) Update(ctx context.Context, annotation *entity.Annotation) error {
	args := m.Called(ctx, annotation)
	return args.Error(0)
}

func (m *AnnotationRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type TagRepository struct {
	mock.Mock
}

func (m *TagRepository) FindByID(ctx context.Context, id int64) (*entity.Tag, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Tag), args.Error(1)
}

func (m *TagRepository) FindByName(ctx context.Context, name string) (*entity.Tag, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Tag), args.Error(1)
}

func (m *TagRepository) FindAll(ctx context.Context) ([]*entity.Tag, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Tag), args.Error(1)
}

func (m *TagRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.Tag, error) {
	args := m.Called(ctx, postID)
	return args.Get(0).([]*entity.Tag), args.Error(1)
}

func (m *TagRepository) FindByPostIDs(ctx context.Context, postIDs []int64) (map[int64][]*entity.Tag, error) {
	args := m.Called(ctx, postIDs)
	return args.Get(0).(map[int64][]*entity.Tag), args.Error(1)
}

func (m *TagRepository) Create(ctx context.Context, tag *entity.Tag) error {
	args := m.Called(ctx, tag)
	return args.Error(0)
}

func (m *TagRepository) Update(ctx context.Context, tag *entity.Tag) error {
	args := m.Called(ctx, tag)
	return args.Error(0)
}

func (m *TagRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type PostTagRepository struct {
	mock.Mock
}

func (m *PostTagRepository) FindByPostID(ctx context.Context, postID int64) ([]*entity.PostTag, error) {
	args := m.Called(ctx, postID)
	return args.Get(0).([]*entity.PostTag), args.Error(1)
}

func (m *PostTagRepository) FindByTagID(ctx context.Context, tagID int64) ([]*entity.PostTag, error) {
	args := m.Called(ctx, tagID)
	return args.Get(0).([]*entity.PostTag), args.Error(1)
}

func (m *PostTagRepository) Create(ctx context.Context, postTag *entity.PostTag) error {
	args := m.Called(ctx, postTag)
	return args.Error(0)
}

func (m *PostTagRepository) Delete(ctx context.Context, postID, tagID int64) error {
	args := m.Called(ctx, postID, tagID)
	return args.Error(0)
}

type TransactionManager struct {
	mock.Mock
}

func (m *TransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	args := m.Called(ctx, fn)
	if fn != nil {
		fn(ctx)
	}
	return args.Error(0)
}