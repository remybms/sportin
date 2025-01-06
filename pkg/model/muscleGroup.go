package model

import (
	"errors"
	"net/http"
)

type MuscleGroupRequest struct {
	Name        string `json:"name"`
	BodyPart    string `json:"body_part"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

type MuscleGroupResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BodyPart    string `json:"body_part"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

func (muscleGroup *MuscleGroupRequest) Bind(r *http.Request) error {
	if muscleGroup.Name == "" {
		return errors.New("name can't be null")
	}
	if muscleGroup.BodyPart == "" {
		return errors.New("body_part can't be null")
	}
	if muscleGroup.Description == "" {
		return errors.New("description can't be null")
	}
	if muscleGroup.Level == "" {
		return errors.New("level can't be null")
	}
	return nil
}
