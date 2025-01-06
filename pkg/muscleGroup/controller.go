package musclegroup

import (
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type MuscleGroup struct {
	*config.Config
}

func New(config *config.Config) *MuscleGroup {
	return &MuscleGroup{config}
}

func (config *MuscleGroup) Create(w http.ResponseWriter, r *http.Request) {
	request := &model.MuscleGroupRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	muscleGroup := &dbmodel.MuscleGroupEntry{Name: request.Name, BodyPart: request.BodyPart, Description: request.Description, Level: request.Level}
	config.MuscleGroupEntryRepository.Create(muscleGroup)
	render.JSON(w, r, config.MuscleGroupEntryRepository.ToModel(muscleGroup))
}

func (config *MuscleGroup) Get(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	muscleGroup, err := config.MuscleGroupEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Muscle group not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, config.MuscleGroupEntryRepository.ToModel(muscleGroup))
}

func (config *MuscleGroup) GetAll(w http.ResponseWriter, r *http.Request) {
	muscleGroups, err := config.MuscleGroupEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching muscle groups", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.MuscleGroupEntryRepository.ToModelList(muscleGroups))
}

func (config *MuscleGroup) Update(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	request := &model.MuscleGroupRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	muscleGroup, err := config.MuscleGroupEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Muscle group not found", http.StatusNotFound)
		return
	}
	muscleGroup.Name = request.Name
	muscleGroup.BodyPart = request.BodyPart
	muscleGroup.Description = request.Description
	muscleGroup.Level = request.Level
	config.MuscleGroupEntryRepository.Update(muscleGroup)
	render.JSON(w, r, config.MuscleGroupEntryRepository.ToModel(muscleGroup))
}

func (config *MuscleGroup) Delete(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	err = config.MuscleGroupEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Muscle group not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, "Muscle group deleted")

}
