package program

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	programConfig := New(configuration)
	router := chi.NewRouter()

	router.Get("/", programConfig.GetAllProgramsHandler)
	router.Get("/{id}", programConfig.GetProgramHandler)
	router.Post("/", programConfig.CreateProgramHandler)
	router.Put("/{id}", programConfig.UpdateProgramHandler)
	router.Delete("/{id}", programConfig.DeleteProgramHandler)
	router.Get("/{id}/exercises", programConfig.GetAllExercicesByProgram)

	return router
}
