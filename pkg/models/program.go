package models

import (
	"errors"
	"net/http"
)

type ProgramRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProgramResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (program *ProgramRequest) Bind(r *http.Request) error {
	if program.Name == "" {
		return errors.New("Name is required")
	}
	if program.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}
