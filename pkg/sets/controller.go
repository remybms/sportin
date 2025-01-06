package sets

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

type SetsConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *SetsConfigurator {
	return &SetsConfigurator{configuration}
}

func (config *SetsConfigurator) CreateSetsHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.SetsRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	setsEntry := &dbmodel.SetsEntry{RPE: req.RPE, RIR: req.RIR, Weight: req.Weight, Work: req.Work, WorkType: req.WorkType, ResistanceBand: req.ResistanceBand, RestTime: req.RestTime}
	config.SetsEntryRepository.Create(setsEntry)

	render.JSON(w, r, config.SetsEntryRepository.ToModel(setsEntry))
}

func (config *SetsConfigurator) GetAllSetsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.SetsEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all sets", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.SetsReponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.SetsEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *SetsConfigurator) GetSetsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.SetsEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieves sets", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.SetsEntryRepository.ToModel(entry))
}

func (config *SetsConfigurator) UpdateSetsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.SetsEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve sets on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, entry)

	entry, err = config.SetsEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update sets on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.SetsEntryRepository.ToModel(entry))
}

func (config *SetsConfigurator) DeleteSetsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	valid, err := config.SetsEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete sets on this id", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Sets does not exist", http.StatusNotFound)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Sets deleted"})
}
