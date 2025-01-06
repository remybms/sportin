package userstats

import (
	"sportin/config"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	UserStatsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Post("/", UserStatsConfigurator.CreateUserStatsHandler)
		r.Get("/", UserStatsConfigurator.GetAllUsersStatsHandler)
		r.Get("/{id}", UserStatsConfigurator.GetUserStatsHandler)
		r.Put("/{id}", UserStatsConfigurator.UpdateUserStatsHandler)
		r.Delete("/{id}", UserStatsConfigurator.DeleteCategoryHandler)
	})
	return router
}
