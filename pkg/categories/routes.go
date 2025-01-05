package categories

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CategoryConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", CategoryConfigurator.addCategoryHandler)
	router.Get("/", CategoryConfigurator.categoriesHandler)
	router.Get("/{id}", CategoryConfigurator.categoryByIdHandler)
	router.Put("/{id}", CategoryConfigurator.editCategoryHandler)
	router.Delete("/{id}", CategoryConfigurator.deleteCategoryHandler)
	return router
}