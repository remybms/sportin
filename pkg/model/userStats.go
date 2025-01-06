package model

import (
	"errors"
	"net/http"
)

type UserStatsRequest struct {
	UserID              int `json:"user_id"`
	Weight              int `json:"weight"`
	Height              int `json:"height"`
	Age                 int `json:"age"`
	ActivityCoefficient int `json:"activity"`
	CaloriesGoal        int `json:"calories_goal"`
	ProteinRatio        int `json:"protein_ratio"`
}

type UserStatsResponse struct {
	ID                  int `json:"id"`
	Weight              int `json:"weight"`
	Height              int `json:"height"`
	Age                 int `json:"age"`
	ActivityCoefficient int `json:"activity"`
	CaloriesGoal        int `json:"calories_goal"`
	ProteinRatio        int `json:"protein_ratio"`
}

func (userStats *UserStatsRequest) Bind(r *http.Request) error {
	if userStats.Weight == 0 {
		return errors.New("Weight is required")
	}
	if userStats.Height == 0 {
		return errors.New("Height is required")
	}
	if userStats.Age == 0 {
		return errors.New("Age is required")
	}
	if userStats.ActivityCoefficient == 0 {
		return errors.New("ActivityCoefficient is required")
	}
	if userStats.CaloriesGoal == 0 {
		return errors.New("CaloriesGoal is required")
	}
	if userStats.ProteinRatio == 0 {
		return errors.New("ProteinRatio is required")
	}
	return nil
}
