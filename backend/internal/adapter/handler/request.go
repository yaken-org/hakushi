package handler

import (
	"errors"
	"strings"

	"github.com/yaken-org/hakushi/internal/usecase"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type CreatePostRequest struct {
	UserAccountID string                   `json:"user_account_id"`
	ImageID       int64                    `json:"image_id"`
	Title         string                   `json:"title"`
	Content       string                   `json:"content"`
	Annotations   []CreateAnnotationRequest `json:"annotations"`
	Tags          []string                 `json:"tags"`
}

type CreateAnnotationRequest struct {
	ProductID   *int64 `json:"product_id"`
	DisplayName string `json:"display_name"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
}

func (r *CreatePostRequest) Validate() error {
	if r.UserAccountID == "" {
		return errors.New("user_account_id is required")
	}
	if r.ImageID == 0 {
		return errors.New("image_id is required")
	}
	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(r.Content) == "" {
		return errors.New("content is required")
	}
	
	for _, annotation := range r.Annotations {
		if strings.TrimSpace(annotation.DisplayName) == "" {
			return errors.New("annotation display_name is required")
		}
		if annotation.X < 0 || annotation.Y < 0 {
			return errors.New("annotation coordinates must be non-negative")
		}
	}
	
	return nil
}

func (r *CreatePostRequest) ToUsecaseInput() usecase.CreatePostInput {
	userAccountID := int64(0)
	if id, err := parseUserAccountID(r.UserAccountID); err == nil {
		userAccountID = id
	}
	
	annotations := make([]usecase.CreateAnnotationInput, len(r.Annotations))
	for i, annotation := range r.Annotations {
		annotations[i] = usecase.CreateAnnotationInput{
			ProductID:   annotation.ProductID,
			DisplayName: annotation.DisplayName,
			X:           annotation.X,
			Y:           annotation.Y,
		}
	}
	
	return usecase.CreatePostInput{
		UserAccountID: userAccountID,
		ImageID:       r.ImageID,
		Title:         strings.TrimSpace(r.Title),
		Content:       strings.TrimSpace(r.Content),
		Annotations:   annotations,
		Tags:          r.Tags,
	}
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *UpdatePostRequest) Validate() error {
	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(r.Content) == "" {
		return errors.New("content is required")
	}
	return nil
}

type CreateUserAccountRequest struct {
	Sub      string `json:"sub"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func (r *CreateUserAccountRequest) Validate() error {
	if strings.TrimSpace(r.Sub) == "" {
		return errors.New("sub is required")
	}
	if strings.TrimSpace(r.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func (r *CreateUserAccountRequest) ToUsecaseInput() usecase.CreateUserAccountInput {
	return usecase.CreateUserAccountInput{
		Sub:      strings.TrimSpace(r.Sub),
		Username: strings.TrimSpace(r.Username),
		Email:    strings.TrimSpace(r.Email),
		Name:     strings.TrimSpace(r.Name),
		ImageURL: strings.TrimSpace(r.ImageURL),
	}
}

type UpdateUserAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func (r *UpdateUserAccountRequest) Validate() error {
	if strings.TrimSpace(r.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func parseUserAccountID(idStr string) (int64, error) {
	// This is a placeholder - in the actual implementation,
	// you might need to handle string-based user account IDs differently
	// For now, we'll assume it's a numeric string
	if idStr == "" {
		return 0, errors.New("empty user account ID")
	}
	// Convert string to int64 if needed
	// For compatibility with existing API
	return 1, nil // Placeholder
}