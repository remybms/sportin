package muscle

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	muscle := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", muscle.Create)
		r.Get("/", muscle.GetAll)
		r.Get("/{id}", muscle.Get)
		r.Put("/{id}", muscle.Update)
		r.Delete("/{id}", muscle.Delete)
	})

	return router
}
