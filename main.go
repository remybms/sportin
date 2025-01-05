package main

import (
	"log"
	"net/http"
	"sportin/config"
	"sportin/pkg/categories"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/api/categories", categories.Routes(configuration))
	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Fatal(http.ListenAndServe(":8080", router))
}
