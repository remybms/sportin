package sets

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	SetConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", SetConfigurator.CreateSetsHandler)
		r.Get("/", SetConfigurator.GetAllSetsHandler)
		r.Get("/{id}", SetConfigurator.GetSetsHandler)
		r.Put("/{id}", SetConfigurator.UpdateSetsHandler)
		r.Delete("/{id}", SetConfigurator.DeleteSetsHandler)
	})
	return router
}
