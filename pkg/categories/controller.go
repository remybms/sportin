package categories

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

type CategoriesConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *CategoriesConfigurator {
	return &CategoriesConfigurator{configuration}
}

// @Summary Create a new category
// @Description Create a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body model.CategoryRequest true "Category object that needs to be created"
// @Success 200 {object} model.CategoryResponse
// @Failure 400 {string} string "Invalid request payload"
// @Router /categories [post]
func (config *CategoriesConfigurator) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CategoryRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryEntry := &dbmodel.CategoryEntry{Name: req.Name, Description: req.Description}
	config.CategoryEntryRepository.Create(categoryEntry)

	render.JSON(w, r, config.CategoryEntryRepository.ToModel(categoryEntry))
}

// @Summary Get all categries
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} []model.CategoryResponse
// @Failure 500 {string} string "Failed to retrieves all categories"
// @Router /categories [get]
func (config *CategoriesConfigurator) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.CategoryEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieves all categories", http.StatusInternalServerError)
		return
	}

	if len(entries) == 0 {
		http.Error(w, "No categories found", http.StatusNotFound)
		return
	}

	responseEntries := make([]*model.CategoryResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.CategoryEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

// @Summary Get a categry
// @Description Get a category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} model.CategoryResponse
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to retrieve category on this id"
// @Router /categories/{id} [get]
func (config *CategoriesConfigurator) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	entry, err := config.CategoryEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve category on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.CategoryEntryRepository.ToModel(entry))
}

// @Summary Update a categry
// @Description Update a category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} model.CategoryResponse
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to update category on this id"
// @Router /categories/{id} [put]
func (config *CategoriesConfigurator) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.CategoryEntryRepository.FindById(id)
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

	entry, err = config.CategoryEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update category on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.CategoryEntryRepository.ToModel(entry))
}

// @Summary Delete category
// @Description Delete category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 500 {string} string "Failed to delete category on this id"
// @Router /categories/{id} [delete]
func (config *CategoriesConfigurator) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	_, err = config.CategoryEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "No category found", http.StatusNotFound)
		return
	}

	valid, err := config.CategoryEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete category on this id", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Category does not exist", http.StatusNotFound)
		return
	}

	render.JSON(w, r, map[string]string{"message": "Category deleted"})
}
