package model

import (
	"errors"
	"net/http"
)

type SetsRequest struct {
	RPE               int    `json:"rpe"`
	RIR               int    `json:"rir"`
	Weight            int    `json:"weight"`
	Work              string `json:"work"`
	WorkType          string `json:"workType"`
	ResistanceBand    string `json:"resistance_band"`
	IntensificationID int    `json:"intensification_id"`
	ProgramExerciseID int    `json:"program_exercise_id"`
	RestTime          int    `json:"rest_time"`
}

type SetsReponse struct {
	ID                int    `json:"id"`
	RPE               int    `json:"rpe"`
	RIR               int    `json:"rir"`
	Weight            int    `json:"weight"`
	Work              string `json:"work"`
	WorkType          string `json:"workType"`
	ResistanceBand    string `json:"resistance_band"`
	IntensificationID int    `json:"intensification_id"`
	ProgramExerciseID int    `json:"program_exercise_id"`
	RestTime          int    `json:"rest_time"`
}

func (sets *SetsRequest) Bind(r *http.Request) error {
	if sets.RPE == 0 {
		return errors.New("RPE is required")
	}
	if sets.RIR == 0 {
		return errors.New("RIR is required")
	}
	if sets.Weight == 0 {
		return errors.New("Weight is required")
	}
	if sets.Work == "" {
		return errors.New("Work is required")
	}
	if sets.WorkType == "" {
		return errors.New("Work Type is required")
	}
	if sets.ResistanceBand == "" {
		return errors.New("Resistance Band is required")
	}
	if sets.IntensificationID == 0 {
		return errors.New("Intensification ID is required")
	}
	if sets.ProgramExerciseID == 0 {
		return errors.New("Program Exercise ID is required")
	}
	if sets.RestTime == 0 {
		return errors.New("Rest Time is required")
	}
	return nil
}
