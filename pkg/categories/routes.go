package categories

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CategoryConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", CategoryConfigurator.CreateCategoryHandler)
		r.Get("/", CategoryConfigurator.GetAllCategoriesHandler)
		r.Get("/{id}", CategoryConfigurator.GetCategoryHandler)
		r.Put("/{id}", CategoryConfigurator.UpdateCategoryHandler)
		r.Delete("/{id}", CategoryConfigurator.DeleteCategoryHandler)
	})
	return router
}
