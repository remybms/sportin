package exercise

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	exercise := New(configuration)
	router := chi.NewRouter()

	router.Post("/", exercise.Create)
	router.Get("/", exercise.GetAll)
	router.Get("/{id}", exercise.Get)
	router.Put("/{id}", exercise.Update)
	router.Delete("/{id}", exercise.Delete)

	return router
}
