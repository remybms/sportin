package exercise

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	exercise := New(configuration)
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", exercise.Create)
		r.Get("/", exercise.GetAll)
		r.Get("/{id}", exercise.Get)
		r.Put("/{id}", exercise.Update)
		r.Delete("/{id}", exercise.Delete)
	})

	return router
}
