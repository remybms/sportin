package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type ExerciseEntry struct {
	gorm.Model
	Name            string `gorm:"column:name"`
	Description     string `gorm:"column:desciption"`
	WeightIncrement int    `gorm:"column:weight_increment"`
	MuscleGroup     MuscleGroupEntry
	MuscleGroupID   int `gorm:"column:muscle_group_id"`
	User            UserEntry
	UserID          int `gorm:"column:user_id"`
}

type ExerciseEntryRepository interface {
	Create(exerciseEntry *ExerciseEntry) (*ExerciseEntry, error)
	FindAll() ([]*ExerciseEntry, error)
	FindById(id int) (*ExerciseEntry, error)
	Update(exerciseEntry *ExerciseEntry) (*ExerciseEntry, error)
	Delete(id int) error
	ToModel(exerciseEntry *ExerciseEntry) *model.ExerciseResponse
	ToModelList(exerciseEntries []*ExerciseEntry) []*model.ExerciseResponse
}

type exerciseEntryRepository struct {
	db *gorm.DB
}

func NewExerciseEntryRepository(db *gorm.DB) ExerciseEntryRepository {
	return &exerciseEntryRepository{db: db}
}

func (r *exerciseEntryRepository) Create(exerciseEntry *ExerciseEntry) (*ExerciseEntry, error) {
	if err := r.db.Create(exerciseEntry).Error; err != nil {
		return nil, err
	}
	return exerciseEntry, nil
}

func (r *exerciseEntryRepository) FindAll() ([]*ExerciseEntry, error) {
	var exerciseEntries []*ExerciseEntry
	if err := r.db.Find(&exerciseEntries).Error; err != nil {
		return nil, err
	}
	return exerciseEntries, nil
}

func (r *exerciseEntryRepository) FindById(id int) (*ExerciseEntry, error) {
	var exerciseEntry ExerciseEntry
	if err := r.db.First(&exerciseEntry, id).Error; err != nil {
		return nil, err
	}
	return &exerciseEntry, nil
}

func (r *exerciseEntryRepository) Update(exerciseEntry *ExerciseEntry) (*ExerciseEntry, error) {
	if err := r.db.Save(exerciseEntry).Error; err != nil {
		return nil, err
	}
	return exerciseEntry, nil
}

func (r *exerciseEntryRepository) Delete(id int) error {
	if err := r.db.Delete(&ExerciseEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *exerciseEntryRepository) ToModel(exerciseEntry *ExerciseEntry) *model.ExerciseResponse {
	return &model.ExerciseResponse{
		ID:              int(exerciseEntry.ID),
		Name:            exerciseEntry.Name,
		Description:     exerciseEntry.Description,
		WeightIncrement: exerciseEntry.WeightIncrement,
		MuscleGroupID:   exerciseEntry.MuscleGroupID,
		UserID:          exerciseEntry.UserID,
	}
}

func (r *exerciseEntryRepository) ToModelList(exerciseEntries []*ExerciseEntry) []*model.ExerciseResponse {
	var modelList []*model.ExerciseResponse
	for _, exerciseEntry := range exerciseEntries {
		modelList = append(modelList, r.ToModel(exerciseEntry))
	}
	return modelList
}
