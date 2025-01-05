package stats

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CategoryConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", CategoryConfigurator.addStatsHandler)
	router.Get("/", CategoryConfigurator.statsHandler)
	router.Get("/{id}", CategoryConfigurator.statByIdHandler)
	router.Put("/{id}", CategoryConfigurator.editStatsHandler)
	router.Delete("/{id}", CategoryConfigurator.deleteStatsHandler)
	return router
}