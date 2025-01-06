package intensifications

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	IntensificationsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", IntensificationsConfigurator.CreateIntensificationHandler)
	router.Get("/", IntensificationsConfigurator.GetAllIntensificationsHandler)
	router.Get("/{id}", IntensificationsConfigurator.GetIntensificationHandler)
	router.Put("/{id}", IntensificationsConfigurator.UpdateIntensificationHandler)
	router.Delete("/{id}", IntensificationsConfigurator.DeleteIntensificationHandler)
	return router
}
