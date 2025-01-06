package model

import (
	"errors"
	"net/http"
)

type ProgramExerciseRequest struct {
	ProgramID  int `json:"program_id"`
	ExerciseID int `json:"exercise_id"`
}

type ProgramExerciseResponse struct {
	ID         int `json:"id"`
	ProgramID  int `json:"program_id"`
	ExerciseID int `json:"exercise_id"`
}

func (programExercise *ProgramExerciseRequest) Bind(r *http.Request) error {
	if programExercise.ProgramID == 0 {
		return errors.New("program_id can't be null")
	}
	if programExercise.ExerciseID == 0 {
		return errors.New("exercise_id can't be null")
	}
	return nil
}
