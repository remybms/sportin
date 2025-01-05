package models

import (
	"net/http"
	"sportin/database/dbmodel"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) Bind(r *http.Request) error {
	panic("unimplemented")
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserResponse(user *dbmodel.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func NewUserResponseList(users []*dbmodel.User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, *NewUserResponse(user))
	}
	return userResponses
}
