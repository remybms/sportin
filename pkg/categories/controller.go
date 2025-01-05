package categories

import (
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CategoriesConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *CategoriesConfigurator {
	return &CategoriesConfigurator{configuration}
}

func categoriesToModel(categories []*dbmodel.Categories) []models.Categories {
	categoriesToModel := &models.Categories{}
	categoriesEdited := []models.Categories{}
	for _, category := range categories {
		categoriesToModel.Name = category.Name
		categoriesToModel.Description = category.Description
		categoriesEdited = append(categoriesEdited, *categoriesToModel)
	}
	return categoriesEdited
}

func (config *CategoriesConfigurator) addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Categories{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addCategory := &dbmodel.Categories{Name: req.Name, Description: req.Description}
	config.CategoriesRepository.Create(addCategory)
	render.JSON(w, r, map[string]string{"success": "New category successfully added"})
}

func (config *CategoriesConfigurator) categoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := config.CategoriesRepository.FindAll()
	categoriesEdited := categoriesToModel(categories)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the categories"})
		return
	}
	render.JSON(w, r, categoriesEdited)
}

func (config *CategoriesConfigurator) categoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "id")
	category, err := config.CategoriesRepository.FindById(categoryId)
	categoryEdited := categoriesToModel(category)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted category"})
		return
	}
	render.JSON(w, r, categoryEdited)
}

func (config *CategoriesConfigurator) editCategoryHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Categories{}
	categoryId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedCategory := &dbmodel.Categories{Name: req.Name, Description: req.Description}
	config.CategoriesRepository.Update(updatedCategory, categoryId)
	render.JSON(w, r, map[string]string{"success": "Category successfully updated"})
}

func (config *CategoriesConfigurator) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "id")
	category, err := config.CategoriesRepository.FindById(categoryId)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted category"})
		return
	}
	config.CategoriesRepository.Delete(category[0])
	render.JSON(w, r, map[string]string{"success": "Category successfully deleted"})
}
