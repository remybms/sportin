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
	"sportin/pkg/sets"
	userstats "sportin/pkg/userStats"
	"sportin/pkg/users"

	_ "sportin/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
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
	router.Mount("/api/v1/sets", sets.Routes(configuration))
	return router
}

// @title Sportin API
// @version 1.0
// @description This is a sample server for a sport application.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	userRepository := configuration.UserRepository

	router := Routes(configuration, userRepository)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
