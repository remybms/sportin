package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type ProgramExerciseEntry struct {
	gorm.Model
	ProgramID  int           `gorm:"column:program_id"`
	Program    ProgramEntry  `gorm:"foreignKey:ProgramID"`
	ExerciseID int           `gorm:"column:exercise_id"`
	Exercise   ExerciseEntry `gorm:"foreignKey:ExerciseID"`
	Sets       *SetsEntry    `gorm:"foreignKey:ProgramExerciceID"`
}

type ProgramExerciseEntryRepository interface {
	Create(programExerciseEntry *ProgramExerciseEntry) (*ProgramExerciseEntry, error)
	FindAll() ([]*ProgramExerciseEntry, error)
	FindById(id int) (*ProgramExerciseEntry, error)
	FindByProgramID(id int) ([]*ExerciseEntry, error)
	Update(programExerciseEntry *ProgramExerciseEntry) (*ProgramExerciseEntry, error)
	Delete(id int) error
	ToModel(programExerciseEntry *ProgramExerciseEntry) *model.ProgramExerciseResponse
	ToModelList(programExerciseEntrys []*ProgramExerciseEntry) []*model.ProgramExerciseResponse
}

type programExerciseEntryRepository struct {
	db *gorm.DB
}

func NewProgramExerciseEntryRepository(db *gorm.DB) ProgramExerciseEntryRepository {
	return &programExerciseEntryRepository{db: db}
}

func (r *programExerciseEntryRepository) Create(programExerciseEntry *ProgramExerciseEntry) (*ProgramExerciseEntry, error) {
	if err := r.db.Create(programExerciseEntry).Error; err != nil {
		return nil, err
	}
	return programExerciseEntry, nil
}

func (r *programExerciseEntryRepository) FindAll() ([]*ProgramExerciseEntry, error) {
	var programExerciseEntrys []*ProgramExerciseEntry
	if err := r.db.Find(&programExerciseEntrys).Error; err != nil {
		return nil, err
	}
	return programExerciseEntrys, nil
}

func (r *programExerciseEntryRepository) FindById(id int) (*ProgramExerciseEntry, error) {
	var programExerciseEntry ProgramExerciseEntry
	if err := r.db.First(&programExerciseEntry, id).Error; err != nil {
		return nil, err
	}
	return &programExerciseEntry, nil
}

func (r *programExerciseEntryRepository) Update(programExerciseEntry *ProgramExerciseEntry) (*ProgramExerciseEntry, error) {
	if err := r.db.Save(programExerciseEntry).Error; err != nil {
		return nil, err
	}
	return programExerciseEntry, nil
}

func (r *programExerciseEntryRepository) Delete(id int) error {
	if err := r.db.Delete(&ProgramExerciseEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *programExerciseEntryRepository) ToModel(programExerciseEntry *ProgramExerciseEntry) *model.ProgramExerciseResponse {
	return &model.ProgramExerciseResponse{
		ID:         int(programExerciseEntry.ID),
		ProgramID:  programExerciseEntry.ProgramID,
		ExerciseID: programExerciseEntry.ExerciseID,
	}
}

func (r *programExerciseEntryRepository) ToModelList(programExerciseEntrys []*ProgramExerciseEntry) []*model.ProgramExerciseResponse {
	var responses []*model.ProgramExerciseResponse
	for _, pe := range programExerciseEntrys {
		responses = append(responses, r.ToModel(pe))
	}
	return responses
}

func (r *programExerciseEntryRepository) FindByProgramID(id int) ([]*ExerciseEntry, error) {
	var programExerciseEntrys []*ProgramExerciseEntry
	if err := r.db.Where("program_id = ?", id).Find(&programExerciseEntrys).Error; err != nil {
		return nil, err
	}
	var exercises []*ExerciseEntry
	for _, pe := range programExerciseEntrys {
		var exercise *ExerciseEntry
		if err := r.db.First(&exercise, pe.ExerciseID).Error; err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}
