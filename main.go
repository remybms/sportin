package main

import (
	"log"
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/categories"
	musclegroup "sportin/pkg/muscleGroup"
	"sportin/pkg/users"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config, userRepository dbmodel.UserRepository) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/api/v1/muscle-group", musclegroup.Routes(configuration))
	router.Mount("/api/v1/users", users.Routes(configuration, userRepository))
	router.Mount("/api/v1/categories", categories.Routes(configuration))
	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	userRepository := configuration.UserRepository

	router := Routes(configuration, userRepository)

	log.Fatal(http.ListenAndServe(":8080", router))
}
