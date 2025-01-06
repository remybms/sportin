package models

import (
	"errors"
	"net/http"
)

type MuscleRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	MuscleGroupID int    `json:"muscle_group_id"`
}

type MuscleResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	MuscleGroupID int    `json:"muscle_group_id"`
}

func (muscle *MuscleRequest) Bind(r *http.Request) error {
	if muscle.Name == "" {
		return errors.New("name can't be null")
	}
	if muscle.Description == "" {
		return errors.New("description can't be null")
	}
	return nil
}
