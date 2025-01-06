package model

import (
	"errors"
	"net/http"
)

type IntensificationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IntensificationResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (intensification *IntensificationRequest) Bind(r *http.Request) error {
	if intensification.Name == "" {
		return errors.New("Name is required")
	}
	if intensification.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}
