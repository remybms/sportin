package users

import (
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/authentification"

	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config, userRepository dbmodel.UserRepository) *chi.Mux {
	router := chi.NewRouter()

	userController := New(config)

	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		r.Get("/", userController.GetUsersHandler)
		r.Get("/{id}", userController.GetUserByIDHandler)
		r.Put("/{id}", userController.UpdateUserHandler)
		r.Delete("/{id}", userController.DeleteUserHandler)
	})

	router.Post("/", userController.CreateUserHandler)
	router.Post("/login", userController.LoginHandler)

	return router
}
