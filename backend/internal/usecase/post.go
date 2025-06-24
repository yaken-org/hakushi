package usecase

import (
	"context"
	"fmt"

	"github.com/yaken-org/hakushi/internal/domain/entity"
	"github.com/yaken-org/hakushi/internal/domain/repository"
)

type PostUsecase struct {
	postRepo       repository.PostRepository
	userRepo       repository.UserAccountRepository
	annotationRepo repository.AnnotationRepository
	tagRepo        repository.TagRepository
	postTagRepo    repository.PostTagRepository
	txManager      TransactionManager
}

func NewPostUsecase(
	postRepo repository.PostRepository,
	userRepo repository.UserAccountRepository,
	annotationRepo repository.AnnotationRepository,
	tagRepo repository.TagRepository,
	postTagRepo repository.PostTagRepository,
	txManager TransactionManager,
) *PostUsecase {
	return &PostUsecase{
		postRepo:       postRepo,
		userRepo:       userRepo,
		annotationRepo: annotationRepo,
		tagRepo:        tagRepo,
		postTagRepo:    postTagRepo,
		txManager:      txManager,
	}
}

func (u *PostUsecase) CreatePost(ctx context.Context, input CreatePostInput) (*PostOutput, error) {
	var result *PostOutput
	
	err := u.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		user, err := u.userRepo.FindByID(ctx, input.UserAccountID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}
		
		post := entity.NewPost(user.ID, input.ImageID, input.Title, input.Content)
		if err := u.postRepo.Create(ctx, post); err != nil {
			return fmt.Errorf("failed to create post: %w", err)
		}
		
		var annotations []*entity.Annotation
		for _, annotationInput := range input.Annotations {
			annotation := entity.NewAnnotation(
				post.ID,
				annotationInput.ProductID,
				annotationInput.DisplayName,
				annotationInput.X,
				annotationInput.Y,
			)
			if err := u.annotationRepo.Create(ctx, annotation); err != nil {
				return fmt.Errorf("failed to create annotation: %w", err)
			}
			annotations = append(annotations, annotation)
		}
		
		var tags []*entity.Tag
		for _, tagName := range input.Tags {
			tag, err := u.tagRepo.FindByName(ctx, tagName)
			if err != nil {
				tag = entity.NewTag(tagName)
				if err := u.tagRepo.Create(ctx, tag); err != nil {
					return fmt.Errorf("failed to create tag: %w", err)
				}
			}
			
			postTag := entity.NewPostTag(post.ID, tag.ID)
			if err := u.postTagRepo.Create(ctx, postTag); err != nil {
				return fmt.Errorf("failed to create post-tag relation: %w", err)
			}
			tags = append(tags, tag)
		}
		
		result = &PostOutput{
			ID:            post.ID,
			UserAccountID: post.UserAccountID,
			ImageID:       post.ImageID,
			Title:         post.Title,
			Content:       post.Content,
			Likes:         post.Likes,
			Annotations:   convertAnnotationsToOutput(annotations),
			Tags:          convertTagsToOutput(tags),
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		}
		
		return nil
	})
	
	return result, err
}

func (u *PostUsecase) GetPostByID(ctx context.Context, id int64) (*PostOutput, error) {
	post, err := u.postRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	
	annotations, err := u.annotationRepo.FindByPostID(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get annotations: %w", err)
	}
	
	tags, err := u.tagRepo.FindByPostID(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	
	return &PostOutput{
		ID:            post.ID,
		UserAccountID: post.UserAccountID,
		ImageID:       post.ImageID,
		Title:         post.Title,
		Content:       post.Content,
		Likes:         post.Likes,
		Annotations:   convertAnnotationsToOutput(annotations),
		Tags:          convertTagsToOutput(tags),
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
	}, nil
}

func (u *PostUsecase) GetAllPosts(ctx context.Context) ([]*PostOutput, error) {
	posts, err := u.postRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	
	if len(posts) == 0 {
		return []*PostOutput{}, nil
	}
	
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	
	annotationsMap, err := u.annotationRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get annotations: %w", err)
	}
	
	tagsMap, err := u.tagRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	
	var result []*PostOutput
	for _, post := range posts {
		annotations := annotationsMap[post.ID]
		tags := tagsMap[post.ID]
		
		output := &PostOutput{
			ID:            post.ID,
			UserAccountID: post.UserAccountID,
			ImageID:       post.ImageID,
			Title:         post.Title,
			Content:       post.Content,
			Likes:         post.Likes,
			Annotations:   convertAnnotationsToOutput(annotations),
			Tags:          convertTagsToOutput(tags),
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		}
		result = append(result, output)
	}
	
	return result, nil
}

func (u *PostUsecase) GetPostsByUserID(ctx context.Context, userID int64) ([]*PostOutput, error) {
	posts, err := u.postRepo.FindByUserAccountID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	
	if len(posts) == 0 {
		return []*PostOutput{}, nil
	}
	
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	
	annotationsMap, err := u.annotationRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get annotations: %w", err)
	}
	
	tagsMap, err := u.tagRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	
	var result []*PostOutput
	for _, post := range posts {
		annotations := annotationsMap[post.ID]
		tags := tagsMap[post.ID]
		
		output := &PostOutput{
			ID:            post.ID,
			UserAccountID: post.UserAccountID,
			ImageID:       post.ImageID,
			Title:         post.Title,
			Content:       post.Content,
			Likes:         post.Likes,
			Annotations:   convertAnnotationsToOutput(annotations),
			Tags:          convertTagsToOutput(tags),
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		}
		result = append(result, output)
	}
	
	return result, nil
}

func (u *PostUsecase) UpdatePost(ctx context.Context, input UpdatePostInput) (*PostOutput, error) {
	post, err := u.postRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	
	post.Title = input.Title
	post.Content = input.Content
	
	if err := u.postRepo.Update(ctx, post); err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}
	
	return u.GetPostByID(ctx, post.ID)
}

func (u *PostUsecase) LikePost(ctx context.Context, postID int64) (int, error) {
	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		return 0, fmt.Errorf("failed to get post: %w", err)
	}
	
	post.IncrementLikes()
	
	if err := u.postRepo.Update(ctx, post); err != nil {
		return 0, fmt.Errorf("failed to update post: %w", err)
	}
	
	return post.Likes, nil
}

func (u *PostUsecase) DeletePost(ctx context.Context, postID int64) error {
	return u.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		postTags, err := u.postTagRepo.FindByPostID(ctx, postID)
		if err != nil {
			return fmt.Errorf("failed to get post tags: %w", err)
		}
		
		for _, postTag := range postTags {
			if err := u.postTagRepo.Delete(ctx, postTag.PostID, postTag.TagID); err != nil {
				return fmt.Errorf("failed to delete post tag: %w", err)
			}
		}
		
		annotations, err := u.annotationRepo.FindByPostID(ctx, postID)
		if err != nil {
			return fmt.Errorf("failed to get annotations: %w", err)
		}
		
		for _, annotation := range annotations {
			if err := u.annotationRepo.Delete(ctx, annotation.ID); err != nil {
				return fmt.Errorf("failed to delete annotation: %w", err)
			}
		}
		
		if err := u.postRepo.Delete(ctx, postID); err != nil {
			return fmt.Errorf("failed to delete post: %w", err)
		}
		
		return nil
	})
}

func (u *PostUsecase) GetPostsOrderByLikes(ctx context.Context) ([]*PostOutput, error) {
	posts, err := u.postRepo.FindOrderByLikes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	
	if len(posts) == 0 {
		return []*PostOutput{}, nil
	}
	
	postIDs := make([]int64, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	
	annotationsMap, err := u.annotationRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get annotations: %w", err)
	}
	
	tagsMap, err := u.tagRepo.FindByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	
	var result []*PostOutput
	for _, post := range posts {
		annotations := annotationsMap[post.ID]
		tags := tagsMap[post.ID]
		
		output := &PostOutput{
			ID:            post.ID,
			UserAccountID: post.UserAccountID,
			ImageID:       post.ImageID,
			Title:         post.Title,
			Content:       post.Content,
			Likes:         post.Likes,
			Annotations:   convertAnnotationsToOutput(annotations),
			Tags:          convertTagsToOutput(tags),
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		}
		result = append(result, output)
	}
	
	return result, nil
}

func convertAnnotationsToOutput(annotations []*entity.Annotation) []AnnotationOutput {
	if annotations == nil {
		return []AnnotationOutput{}
	}
	
	result := make([]AnnotationOutput, len(annotations))
	for i, annotation := range annotations {
		result[i] = AnnotationOutput{
			ID:          annotation.ID,
			PostID:      annotation.PostID,
			ProductID:   annotation.ProductID,
			DisplayName: annotation.DisplayName,
			X:           annotation.X,
			Y:           annotation.Y,
			CreatedAt:   annotation.CreatedAt,
			UpdatedAt:   annotation.UpdatedAt,
		}
	}
	return result
}

func convertTagsToOutput(tags []*entity.Tag) []TagOutput {
	if tags == nil {
		return []TagOutput{}
	}
	
	result := make([]TagOutput, len(tags))
	for i, tag := range tags {
		result[i] = TagOutput{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
		}
	}
	return result
}