package program

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

type ProgramConfig struct {
	*config.Config
}

func New(configuration *config.Config) *ProgramConfig {
	return &ProgramConfig{configuration}
}

// @Summary Create a new program
// @Description Create a new program
// @Tags Program
// @Accept json
// @Produce json
// @Param program body model.ProgramRequest true "Program object that needs to be created"
// @Success 200 {object} model.ProgramResponse
// @Router /programs [post]
func (config *ProgramConfig) CreateProgramHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.ProgramRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	programEntry := &dbmodel.ProgramEntry{UserID: req.UserID, CategoryID: req.CategoryID, Name: req.Name, Description: req.Description}
	config.ProgramEntryRepository.Create(programEntry)

	render.JSON(w, r, config.ProgramEntryRepository.ToModel(programEntry))
}

// @Summary Get all programs
// @Description Get all programs
// @Tags Program
// @Accept json
// @Produce json
// @Success 200 {object} []model.ProgramResponse
// @Router /programs [get]
func (config *ProgramConfig) GetAllProgramsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.ProgramEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all programs", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.ProgramResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.ProgramEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

// @Summary Get program
// @Description Get program
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {object} model.ProgramResponse
// @Router /programs/{id} [get]
func (config *ProgramConfig) GetProgramHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.ProgramEntryRepository.FindByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve program on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.ProgramEntryRepository.ToModel(entry))
}

// @Summary Update a program
// @Description Update a program
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Param program body model.ProgramRequest true "Program object that needs to be updated"
// @Success 200 {object} model.ProgramResponse
// @Router /programs/{id} [put]
func (config *ProgramConfig) UpdateProgramHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.ProgramEntryRepository.FindByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve program on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, entry)

	entry, err = config.ProgramEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update program on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.ProgramEntryRepository.ToModel(entry))
}

// @Summary Delete a program
// @Description Delete a program
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {object} string
// @Router /programs/{id} [delete]
func (config *ProgramConfig) DeleteProgramHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	valid, err := config.ProgramEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete program on this id", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Program does not exist", http.StatusNotFound)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Program deleted"})
}
