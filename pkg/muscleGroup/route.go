package musclegroup

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	muscleGroup := New(configuration)
	router := chi.NewRouter()

	router.Post("/", muscleGroup.Create)
	router.Get("/", muscleGroup.GetAll)
	router.Get("/{id}", muscleGroup.Get)
	router.Put("/{id}", muscleGroup.Update)
	router.Delete("/{id}", muscleGroup.Delete)

	return router
}
