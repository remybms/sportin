package userstats

import (
	"encoding/json"
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/helper"
	"sportin/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserStatsConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *UserStatsConfigurator {
	return &UserStatsConfigurator{configuration}
}

func (config *UserStatsConfigurator) CreateUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.UserStatsRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userStatsEntry := &dbmodel.UserStatsEntry{UserID: req.UserID, Weight: req.Weight, Height: req.Height, Age: req.Age, ActivityCoefficient: req.ActivityCoefficient, CaloriesGoal: req.CaloriesGoal, ProteinRatio: req.ProteinRatio}
	config.UserStatsRepository.Create(userStatsEntry)

	render.JSON(w, r, config.UserStatsRepository.ToModel(userStatsEntry))
}

func (config *UserStatsConfigurator) GetAllUsersStatsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.UserStatsRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all users stats", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.UserStatsResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.UserStatsRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *UserStatsConfigurator) GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.UserStatsRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve user stats on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.UserStatsRepository.ToModel(entry))
}

func (config *UserStatsConfigurator) UpdateUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.UserStatsRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve user stats on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, entry)

	entry, err = config.UserStatsRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update category on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.UserStatsRepository.ToModel(entry))
}

func (config *UserStatsConfigurator) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	valid, err := config.UserStatsRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete user stats on this id", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "User stats does not exist", http.StatusNotFound)
		return
	}
	render.JSON(w, r, map[string]string{"message": "User stats deleted"})
}
