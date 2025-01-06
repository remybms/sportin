package main

import (
	"log"
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/categories"
	"sportin/pkg/exercise"
	"sportin/pkg/intensifications"
	"sportin/pkg/muscle"
	musclegroup "sportin/pkg/muscleGroup"
	"sportin/pkg/program"
	programexercise "sportin/pkg/programExercise"
	userstats "sportin/pkg/userStats"
	"sportin/pkg/users"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config, userRepository dbmodel.UserRepository) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/api/v1/muscle-group", musclegroup.Routes(configuration))
	router.Mount("/api/v1/users", users.Routes(configuration, userRepository))
	router.Mount("/api/v1/users_stats", userstats.Routes(configuration))
	router.Mount("/api/v1/categories", categories.Routes(configuration))
	router.Mount("/api/v1/muscle", muscle.Routes(configuration))
	router.Mount("/api/v1/programs", program.Routes(configuration))
	router.Mount("/api/v1/exercises", exercise.Routes(configuration))
	router.Mount("/api/v1/intensifications", intensifications.Routes(configuration))
	router.Mount("/api/v1/program_exercises", programexercise.Routes(configuration))
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
