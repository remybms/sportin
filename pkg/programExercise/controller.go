package programExercise

import (
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ProgramExercise struct {
	*config.Config
}

func New(config *config.Config) *ProgramExercise {
	return &ProgramExercise{config}
}

// @Summary Create a new program exercise
// @Description Create a new program exercise
// @Tags ProgramExercise
// @Accept json
// @Produce json
// @Param programExercise body model.ProgramExerciseRequest true "Program Exercise object that needs to be created"
// @Success 200 {object} model.ProgramExerciseResponse
// @Failure 400 {string} string err.Error()
// @Failure 500 {string} string "Failed to create program exercise"
// @Router /programExercise [post]
func (config *ProgramExercise) Create(w http.ResponseWriter, r *http.Request) {
	req := &model.ProgramExerciseRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	programExercise := &dbmodel.ProgramExerciseEntry{ProgramID: req.ProgramID, ExerciseID: req.ExerciseID}
	_, err := config.ProgramExerciseEntryRepository.Create(programExercise)
	if err != nil {
		http.Error(w, "Failed to create program exercise", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModel(programExercise))
}

// @Summary Get program exercise
// @Description Get program exercise
// @Tags ProgramExercise
// @Accept json
// @Produce json
// @Param id path int true "Program Exercise ID"
// @Success 200 {object} model.ProgramExerciseResponse
// @Failure 400 {string} string "Missing or invalid id parameter"
// @Failure 404 {string} string "Program exercise not found"
// @Router /programExercise/{id} [get]
func (config *ProgramExercise) Get(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	programExercise, err := config.ProgramExerciseEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Program exercise not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModel(programExercise))
}

// @Summary Get all program exercises
// @Description Get all program exercises
// @Tags ProgramExercise
// @Accept json
// @Produce json
// @Success 200 {object} []model.ProgramExerciseResponse
// @Failure 500 {string} string "Error fetching program exercises"
// @Router /programExercise [get]
func (config *ProgramExercise) GetAll(w http.ResponseWriter, r *http.Request) {
	programExercises, err := config.ProgramExerciseEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching program exercises", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModelList(programExercises))
}

// @Summary Update program exercise
// @Description Update program exercise
// @Tags ProgramExercise
// @Accept json
// @Produce json
// @Param id path int true "Program Exercise ID"
// @Param programExercise body model.ProgramExerciseRequest true "Program Exercise object that needs to be updated"
// @Success 200 {object} model.ProgramExerciseResponse
// @Failure 400 {string} string "Missing or invalid id parameter"
// @Failure 404 {string} string "Program exercise not found"
// @Router /programExercise/{id} [put]
func (config *ProgramExercise) Update(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	request := &model.ProgramExerciseRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	programExercise := &dbmodel.ProgramExerciseEntry{ProgramID: request.ProgramID, ExerciseID: request.ExerciseID}
	config.ProgramExerciseEntryRepository.Update(programExercise)
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModel(programExercise))
}

// @Summary Delete program exercise
// @Description Delete program exercise
// @Tags ProgramExercise
// @Accept json
// @Produce json
// @Param id path int true "Program Exercise ID"
// @Success 200 {string} string "Program exercise deleted"
// @Failure 400 {string} string "Missing id parameter"
// @Failure 500 {string} string "Failed to delete program exercise"
// @Router /programExercise/{id} [delete]
func (config *ProgramExercise) Delete(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	err = config.ProgramExerciseEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete program exercise", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, "Program exercise deleted")
}
