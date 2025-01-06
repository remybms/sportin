package userstats

import (
	"sportin/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	UserStatsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", UserStatsConfigurator.CreateUserStatsHandler)
	router.Get("/", UserStatsConfigurator.GetAllUsersStatsHandler)
	router.Get("/{id}", UserStatsConfigurator.GetUserStatsHandler)
	router.Put("/{id}", UserStatsConfigurator.UpdateUserStatsHandler)
	router.Delete("/{id}", UserStatsConfigurator.DeleteCategoryHandler)
	return router
}
