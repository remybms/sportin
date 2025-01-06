package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type UserStatsEntry struct {
	gorm.Model
	Weight              int       `gorm:"column:weight"`
	Height              int       `gorm:"column:height"`
	Age                 int       `gorm:"column:age"`
	ActivityCoefficient int       `gorm:"column:activity"`
	CaloriesGoal        int       `gorm:"column:calories_goal"`
	ProteinRatio        int       `gorm:"column:protein_ratio"`
	UserID              int       `gorm:"column:user_id"`
	User                UserEntry `gorm:"foreignKey:UserID"`
}

type UserStatsRepository interface {
	Create(userStatsEntry *UserStatsEntry) (*UserStatsEntry, error)
	FindAll() ([]*UserStatsEntry, error)
	FindById(id int) (*UserStatsEntry, error)
	Delete(id int) (bool, error)
	Update(userStatsEntry *UserStatsEntry) (*UserStatsEntry, error)
	ToModel(entry *UserStatsEntry) *model.UserStatsResponse
}

type userStatsRepository struct {
	db *gorm.DB
}

func NewUserStatsRepository(db *gorm.DB) UserStatsRepository {
	return &userStatsRepository{db: db}
}

func (r *userStatsRepository) Create(userStatsEntry *UserStatsEntry) (*UserStatsEntry, error) {
	if err := r.db.Create(userStatsEntry).Error; err != nil {
		return nil, err
	}
	return userStatsEntry, nil
}

func (r *userStatsRepository) FindAll() ([]*UserStatsEntry, error) {
	var userStatsEntries []*UserStatsEntry
	if err := r.db.Find(&userStatsEntries).Error; err != nil {
		return nil, err
	}
	return userStatsEntries, nil
}

func (r *userStatsRepository) FindById(id int) (*UserStatsEntry, error) {
	var userStatsEntry *UserStatsEntry
	if err := r.db.First(&userStatsEntry, id).Error; err != nil {
		return nil, err
	}
	return userStatsEntry, nil
}

func (r *userStatsRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&UserStatsEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *userStatsRepository) Update(userStatsEntry *UserStatsEntry) (*UserStatsEntry, error) {
	if err := r.db.Where("id = ?", userStatsEntry.ID).Updates(userStatsEntry).Error; err != nil {
		return nil, err
	}
	return userStatsEntry, nil
}

func (r *userStatsRepository) ToModel(entry *UserStatsEntry) *model.UserStatsResponse {
	return &model.UserStatsResponse{
		ID:                  int(entry.ID),
		Weight:              entry.Weight,
		Height:              entry.Height,
		Age:                 entry.Age,
		ActivityCoefficient: entry.ActivityCoefficient,
		CaloriesGoal:        entry.CaloriesGoal,
		ProteinRatio:        entry.ProteinRatio,
	}
}
