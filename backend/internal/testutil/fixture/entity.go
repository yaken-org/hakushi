package fixture

import (
	"time"

	"github.com/yaken-org/hakushi/internal/domain/entity"
)

func NewPost(id int64, userAccountID int64, imageID int64, title, content string) *entity.Post {
	now := time.Now()
	return &entity.Post{
		ID:            id,
		UserAccountID: userAccountID,
		ImageID:       imageID,
		Title:         title,
		Content:       content,
		Likes:         0,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func NewUserAccount(id int64, sub, username, email, name, imageURL string) *entity.UserAccount {
	now := time.Now()
	return &entity.UserAccount{
		ID:        id,
		Sub:       sub,
		Username:  username,
		Email:     email,
		Name:      name,
		ImageURL:  imageURL,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewAnnotation(id, postID int64, productID *int64, displayName string, x, y int) *entity.Annotation {
	now := time.Now()
	return &entity.Annotation{
		ID:          id,
		PostID:      postID,
		ProductID:   productID,
		DisplayName: displayName,
		X:           x,
		Y:           y,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func NewTag(id int64, name string) *entity.Tag {
	now := time.Now()
	return &entity.Tag{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewPostTag(id, postID, tagID int64) *entity.PostTag {
	return &entity.PostTag{
		ID:        id,
		PostID:    postID,
		TagID:     tagID,
		CreatedAt: time.Now(),
	}
}