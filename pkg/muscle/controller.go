package muscle

import (
	"encoding/json"
	"log"
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/helper"
	"sportin/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Muscle struct {
	*config.Config
}

func New(config *config.Config) *Muscle {
	return &Muscle{config}
}

// @Summary Create a new muscle
// @Description Create a new muscle
// @Tags Muscle
// @Accept json
// @Produce json
// @Param muscle body model.MuscleRequest true "Muscle object that needs to be created"
// @Success 200 {object} model.MuscleResponse
// @Router /muscle [post]
func (config *Muscle) Create(w http.ResponseWriter, r *http.Request) {
	request := &model.MuscleRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if request.MuscleGroupID < 1 {
		http.Error(w, "Invalid Muscle Group ID", http.StatusBadRequest)
		return
	}
	_, err := config.MuscleGroupEntryRepository.FindById(request.MuscleGroupID)
	if err != nil {
		http.Error(w, "Muscle Group not found", http.StatusNotFound)
	}
	muscle := &dbmodel.MuscleEntry{Name: request.Name, Description: request.Description, MuscleGroupID: request.MuscleGroupID}
	config.MuscleEntryRepository.Create(muscle)
	render.JSON(w, r, config.MuscleEntryRepository.ToModel(muscle))
}

// @Summary Get muscle
// @Description Get muscle
// @Tags Muscle
// @Accept json
// @Produce json
// @Param id path int true "Muscle ID"
// @Success 200 {object} model.MuscleResponse
// @Router /muscle/{id} [get]
func (config *Muscle) Get(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	muscle, err := config.MuscleEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Muscle  not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, config.MuscleEntryRepository.ToModel(muscle))
}

// @Summary Get all muscles
// @Description Get all muscles
// @Tags Muscle
// @Accept json
// @Produce json
// @Success 200 {object} []model.MuscleResponse
// @Router /muscle [get]
func (config *Muscle) GetAll(w http.ResponseWriter, r *http.Request) {
	muscles, err := config.MuscleEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching muscle s", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, config.MuscleEntryRepository.ToModelList(muscles))
}

// @Summary Update muscle
// @Description Update muscle
// @Tags Muscle
// @Accept json
// @Produce json
// @Param id path int true "Muscle ID"
// @Param muscle body model.MuscleRequest true "Muscle object that needs to be updated"
// @Success 200 {object} model.MuscleResponse
// @Router /muscle/{id} [put]
func (config *Muscle) Update(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	muscle, err := config.MuscleEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Muscle  not found", http.StatusNotFound)
		return
	}
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	helper.ApplyChanges(data, muscle)
	updatedMuscle, err := config.MuscleEntryRepository.Update(muscle)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}
	log.Println(updatedMuscle)
	log.Println(data)
	render.JSON(w, r, config.MuscleEntryRepository.ToModel(updatedMuscle))
}

// @Summary Delete muscle
// @Description Delete muscle
// @Tags Muscle
// @Accept json
// @Produce json
// @Param id path int true "Muscle ID"
// @Success 200 {object} string
// @Router /muscle/{id} [delete]
func (config *Muscle) Delete(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	if strId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}
	err = config.MuscleEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Muscle  not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, "Muscle  deleted")

}
