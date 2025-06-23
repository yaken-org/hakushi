package repository

import "errors"

var (
	ErrPostNotFound        = errors.New("post not found")
	ErrUserAccountNotFound = errors.New("user account not found")
	ErrAnnotationNotFound  = errors.New("annotation not found")
	ErrTagNotFound         = errors.New("tag not found")
	ErrProductNotFound     = errors.New("product not found")
	ErrPostTagNotFound     = errors.New("post tag not found")
)