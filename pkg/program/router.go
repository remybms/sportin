package program

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	programConfig := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Get("/", programConfig.GetAllProgramsHandler)
		r.Get("/{id}", programConfig.GetProgramHandler)
		r.Post("/", programConfig.CreateProgramHandler)
		r.Put("/{id}", programConfig.UpdateProgramHandler)
		r.Delete("/{id}", programConfig.DeleteProgramHandler)
	})

	return router
}
