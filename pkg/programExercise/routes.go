package programExercise

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	programExercise := New(configuration)
	router := chi.NewRouter()

	router.Post("/", programExercise.Create)
	router.Get("/", programExercise.GetAll)
	router.Get("/{id}", programExercise.Get)
	router.Put("/{id}", programExercise.Update)
	router.Delete("/{id}", programExercise.Delete)

	return router
}
