package exercise

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

type exerciseConfig struct {
	*config.Config
}

func New(config *config.Config) *exerciseConfig {
	return &exerciseConfig{config}
}

// @Summary Create a new exercise
// @Description Delete a new exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param exercise body model.ExerciseRequest true "Exercise object that needs to be created"
// @Success 200 {object} model.ExerciseResponse
// @Router /exercise [post]
func (config *exerciseConfig) Create(w http.ResponseWriter, r *http.Request) {
	req := &model.ExerciseRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ExerciseEntry := &dbmodel.ExerciseEntry{UserID: req.UserID, Name: req.Name, Description: req.Description, MuscleGroupID: req.MuscleGroupID, WeightIncrement: req.WeightIncrement}
	config.ExerciseEntryRepository.Create(ExerciseEntry)
	render.JSON(w, r, config.ExerciseEntryRepository.ToModel(ExerciseEntry))
}

func (config *exerciseConfig) GetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := config.ExerciseEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all programs", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.ExerciseResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.ExerciseEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

// @Summary Get exercise
// @Description Get exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} []model.ExerciseResponse
// @Router /exercise [get]
func (config *exerciseConfig) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.ExerciseEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieves all programs", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.ExerciseEntryRepository.ToModel(entry))
}

// @Summary Update a new exercise
// @Description Update a new exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Param exercise body model.ExerciseRequest true "Exercise object that needs to be updated"
// @Success 200 {object} model.ExerciseResponse
// @Router /exercise/{id} [put]
func (config *exerciseConfig) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.ExerciseEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve category on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, entry)

	entry, err = config.ExerciseEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update category on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.ExerciseEntryRepository.ToModel(entry))
}

// @Summary Delete a new exercise
// @Description Delete a new exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} string
// @Router /exercise/{id} [delete]
func (config *exerciseConfig) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	err = config.ExerciseEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete program on this id", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, "Program deleted")
}
