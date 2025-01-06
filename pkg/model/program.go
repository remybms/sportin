package model

import (
	"errors"
	"net/http"
)

type ProgramRequest struct {
	UserID      int    `json:"user_id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProgramResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (program *ProgramRequest) Bind(r *http.Request) error {
	if program.UserID < 0 {
		return errors.New("Invalid user id")
	}
	if program.CategoryID < 0 {
		return errors.New("Invalid category id")
	}
	if program.Name == "" {
		return errors.New("Name is required")
	}
	if program.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}
