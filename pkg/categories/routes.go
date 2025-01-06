package categories

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CategoryConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", CategoryConfigurator.CreateCategoryHandler)
	router.Get("/", CategoryConfigurator.GetAllCategoriesHandler)
	router.Get("/{id}", CategoryConfigurator.GetCategoryHandler)
	router.Put("/{id}", CategoryConfigurator.UpdateCategoryHandler)
	router.Delete("/{id}", CategoryConfigurator.DeleteCategoryHandler)
	return router
}
