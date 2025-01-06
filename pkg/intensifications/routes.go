package intensifications

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	IntensificationsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", IntensificationsConfigurator.CreateIntensificationHandler)
		r.Get("/", IntensificationsConfigurator.GetAllIntensificationsHandler)
		r.Get("/{id}", IntensificationsConfigurator.GetIntensificationHandler)
		r.Put("/{id}", IntensificationsConfigurator.UpdateIntensificationHandler)
		r.Delete("/{id}", IntensificationsConfigurator.DeleteIntensificationHandler)
	})
	return router
}
