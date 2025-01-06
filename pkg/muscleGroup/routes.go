package musclegroup

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	muscleGroup := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", muscleGroup.Create)
		r.Get("/", muscleGroup.GetAll)
		r.Get("/{id}", muscleGroup.Get)
		r.Put("/{id}", muscleGroup.Update)
		r.Delete("/{id}", muscleGroup.Delete)
	})

	return router
}
