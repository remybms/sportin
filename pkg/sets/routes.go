package sets

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	SetConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", SetConfigurator.CreateSetsHandler)
	router.Get("/", SetConfigurator.GetAllSetsHandler)
	router.Get("/{id}", SetConfigurator.GetSetsHandler)
	router.Put("/{id}", SetConfigurator.UpdateSetsHandler)
	router.Delete("/{id}", SetConfigurator.DeleteSetsHandler)
	return router
}
