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
