package programExercise

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	programExercise := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", programExercise.Create)
		r.Get("/", programExercise.GetAll)
		r.Get("/{id}", programExercise.Get)
		r.Put("/{id}", programExercise.Update)
		r.Delete("/{id}", programExercise.Delete)
	})

	return router
}
