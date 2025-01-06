package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type UserEntry struct {
	gorm.Model
	Username  string           `gorm:"column:username"`
	Email     string           `gorm:"column:email"`
	Password  string           `gorm:"column:password"`
	UserStats *UserStatsEntry  `gorm:"foreignKey:UserID"`
	Programs  []*ProgramEntry  `gorm:"foreignKey:UserID"`
	Exercises []*ExerciseEntry `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	Create(user *UserEntry) (*UserEntry, error)
	FindAll() ([]*UserEntry, error)
	FindByID(id int) (*UserEntry, error)
	FindByEmail(email string) (*UserEntry, error)
	Update(user *UserEntry) (*UserEntry, error)
	Delete(id int) error
	ToModel(user *UserEntry) *model.UserResponse
	ToModelList(users []*UserEntry) []*model.UserResponse
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *UserEntry) (*UserEntry, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*UserEntry, error) {
	var users []*UserEntry
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id int) (*UserEntry, error) {
	var user UserEntry
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*UserEntry, error) {
	var user UserEntry
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *UserEntry) (*UserEntry, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&UserEntry{}, id).Error
}

func (r *userRepository) ToModel(user *UserEntry) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (r *userRepository) ToModelList(users []*UserEntry) []*model.UserResponse {
	var userResponses []*model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, r.ToModel(user))
	}
	return userResponses
}
