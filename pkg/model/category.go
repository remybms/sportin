package model

import (
	"errors"
	"net/http"
)

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (category *CategoryRequest) Bind(r *http.Request) error {
	if category.Name == "" {
		return errors.New("Name is required")
	}
	if category.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}
