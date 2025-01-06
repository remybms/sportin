package dbmodel

import (
	"sportin/pkg/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Stats    []*Stats `gorm:"foreignKey:UserId"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uint) error
	ToModel(user *User) *models.UserResponse
	ToModelList(users []*User) []*models.UserResponse
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*User, error) {
	var users []*User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *userRepository) ToModel(user *User) *models.UserResponse {
	return &models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (r *userRepository) ToModelList(users []*User) []*models.UserResponse {
	var userResponses []*models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, r.ToModel(user))
	}
	return userResponses
}
