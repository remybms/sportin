package users

import (
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	*config.Config
	UserRepository dbmodel.UserRepository
}

func New(config *config.Config, userRepository dbmodel.UserRepository) *UserController {
	return &UserController{
		Config:         config,
		UserRepository: userRepository,
	}
}

func (controller *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to process password"})
		return
	}

	user := &dbmodel.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	newUser, err := controller.UserRepository.Create(user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}

	render.JSON(w, r, models.NewUserResponse(newUser))
}
func (controller *UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := controller.UserRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve users"})
		return
	}

	render.JSON(w, r, models.NewUserResponseList(users))
}

func (controller *UserController) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	user, err := controller.UserRepository.FindByID(uint(userID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	render.JSON(w, r, models.NewUserResponse(user))
}

func (controller *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	user, err := controller.UserRepository.FindByID(uint(userID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	user.Username = req.Username
	user.Email = req.Email
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Failed to process password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	updatedUser, err := controller.UserRepository.Update(user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	render.JSON(w, r, models.NewUserResponse(updatedUser))
}

func (controller *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	if err := controller.UserRepository.Delete(uint(userID)); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User deleted successfully"})
}
