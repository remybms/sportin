package models

import "net/http"

type UserStats struct {
	UserId              int `json:"user_id"`
	Weight              int `json:"user_weight"`
	Height              int `json:"user_height"`
	Age                 int `json:"user_age"`
	ActivityCoefficient int `json:"user_activity"`
	CaloriesGoal        int `json:"user_calories_goal"`
	ProteinRatio        int `json:"user_protein_ratio"`
}

func (c *UserStats) Bind(r *http.Request) error {
	return nil
}
