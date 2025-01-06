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

func (config *ProgramExercise) Create(w http.ResponseWriter, r *http.Request) {
	req := &model.ProgramExerciseRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	programExercise := &dbmodel.ProgramExerciseEntry{ProgramID: req.ProgramID, ExerciseID: req.ExerciseID}
	config.ProgramExerciseEntryRepository.Create(programExercise)
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModel(programExercise))
}

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

func (config *ProgramExercise) GetAll(w http.ResponseWriter, r *http.Request) {
	programExercises, err := config.ProgramExerciseEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching program exercises", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.ProgramExerciseEntryRepository.ToModelList(programExercises))
}

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
