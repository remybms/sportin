package users

import (
	"encoding/json"
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/helper"
	"sportin/pkg/authentification"
	"sportin/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	*config.Config
}

func New(config *config.Config) *UserConfig {
	return &UserConfig{config}
}

func (config *UserConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to process password"})
		return
	}

	user := &dbmodel.UserEntry{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	newUser, err := config.UserRepository.Create(user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}

	render.JSON(w, r, config.UserRepository.ToModel(newUser))
}

func (config *UserConfig) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := config.UserRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve users"})
		return
	}

	render.JSON(w, r, config.UserRepository.ToModelList(users))
}

func (config *UserConfig) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	user, err := config.UserRepository.FindByID(userID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	render.JSON(w, r, config.UserRepository.ToModel(user))
}

func (config *UserConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.UserRequest{}
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

	var data map[string]interface{}

	user, err := config.UserRepository.FindByID(userID)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	helper.ApplyChanges(data, user)

	_, err = config.UserRepository.Update(user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	render.JSON(w, r, config.UserRepository.ToModel(user))
}

func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	if err := config.UserRepository.Delete(userID); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User deleted successfully"})
}

func (config *UserConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userEntry, err := config.UserRepository.FindByEmail(payload.Email)
	if err != nil {
		http.Error(w, "Unknow user", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEntry.Password), []byte(payload.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := authentification.GenerateJWTToken("your_secret_key", userEntry)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, (map[string]string{"token": token}))
}
