package intensifications

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

type IntensificationsConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *IntensificationsConfigurator {
	return &IntensificationsConfigurator{configuration}
}

// @Summary Create a new Intensification
// @Description Create a new Intensification
// @Tags Intensification
// @Accept json
// @Produce json
// @Param exercise body model.IntensificationRequest true "Exercise object that needs to be created"
// @Success 200 {object} model.IntensificationResponse
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to create intensification"
// @Router /intensifications [post]
func (config *IntensificationsConfigurator) CreateIntensificationHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.IntensificationRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intensificationEntry := &dbmodel.IntensificationEntry{Name: req.Name, Description: req.Description}
	config.IntensificationEntryRepository.Create(intensificationEntry)

	render.JSON(w, r, config.IntensificationEntryRepository.ToModel(intensificationEntry))
}

// @Summary Get all Intensification
// @Description Get all Intensification
// @Tags Intensification
// @Accept json
// @Produce json
// @Param Category body model.IntensificationRequest true "Intensification object that needs to be created"
// @Success 200 {object} model.IntensificationResponse
// @Failure 500 {string} string "Failed to retrieves all intensifications"
// @Router /intensifications [get]
func (config *IntensificationsConfigurator) GetAllIntensificationsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.IntensificationEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all intensifications", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.IntensificationResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.IntensificationEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

// @Summary Get Intensification
// @Description Get Intensification
// @Tags Intensification
// @Accept json
// @Produce json
// @Param id path int true "Intensification ID"
// @Success 200 {object} model.IntensificationResponse
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to retrieves intensification"
// @Router /intensifications/{id} [get]
func (config *IntensificationsConfigurator) GetIntensificationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.IntensificationEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieves intensification", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.IntensificationEntryRepository.ToModel(entry))
}

// @Summary Update Intensification
// @Description Update Intensification
// @Tags Intensification
// @Accept json
// @Produce json
// @Param id path int true "Intensification ID"
// @Param intensification body model.IntensificationRequest true "Intensification object that needs to be updated"
// @Success 200 {object} model.IntensificationResponse
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to update intensification"
// @Router /intensifications/{id} [put]
func (config *IntensificationsConfigurator) UpdateIntensificationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.IntensificationEntryRepository.FindById(id)
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

	entry, err = config.IntensificationEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update category on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.IntensificationEntryRepository.ToModel(entry))
}

// @Summary Delete Intensification
// @Description Delete Intensification
// @Tags Intensification
// @Accept json
// @Produce json
// @Param id path int true "Intensification ID"
// @Success 200 {object} string
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to delete intensification on this id"
// @Failure 404 {string} string "Intensification does not exist"
// @Router /intensifications/{id} [delete]
func (config *IntensificationsConfigurator) DeleteIntensificationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	valid, err := config.IntensificationEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete intensification on this id", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Intensification does not exist", http.StatusNotFound)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Intensification deleted"})
}
