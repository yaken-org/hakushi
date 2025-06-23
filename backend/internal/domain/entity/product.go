package entity

import (
	"time"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int
	ImageURL    string
	URL         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name, description string, price int, imageURL, url string) *Product {
	now := time.Now()
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL,
		URL:         url,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}