package users

import (
	"sportin/config"
	"sportin/database/dbmodel"

	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config, userRepository dbmodel.UserRepository) *chi.Mux {
	router := chi.NewRouter()

	userController := New(config)

	router.Post("/", userController.CreateUserHandler)
	router.Get("/", userController.GetUsersHandler)
	router.Get("/{id}", userController.GetUserByIDHandler)
	router.Put("/{id}", userController.UpdateUserHandler)
	router.Delete("/{id}", userController.DeleteUserHandler)

	return router
}
