package muscle

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	muscle := New(configuration)
	router := chi.NewRouter()

	router.Post("/", muscle.Create)
	router.Get("/", muscle.GetAll)
	router.Get("/{id}", muscle.Get)
	router.Put("/{id}", muscle.Update)
	router.Delete("/{id}", muscle.Delete)

	return router
}
