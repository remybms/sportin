package model

import (
	"errors"
	"net/http"
)

type ExerciseRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	WeightIncrement int    `json:"weight_increment"`
	MuscleGroupID   int    `json:"muscle_group_id"`
	UserID          int    `json:"user_id"`
}

type ExerciseResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	WeightIncrement int    `json:"weight_increment"`
	MuscleGroupID   int    `json:"muscle_group_id"`
	UserID          int    `json:"user_id"`
}

func (exercise *ExerciseRequest) Bind(r *http.Request) error {
	if exercise.Name == "" {
		return errors.New("name can't be null")
	}
	if exercise.Description == "" {
		return errors.New("description can't be null")
	}
	if exercise.WeightIncrement == 0 {
		return errors.New("weight increment can't be null")
	}
	if exercise.MuscleGroupID == 0 {
		return errors.New("muscle group id can't be null")
	}
	if exercise.UserID == 0 {
		return errors.New("user id can't be null")
	}
	return nil
}
