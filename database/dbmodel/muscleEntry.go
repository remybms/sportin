package dbmodel

import (
	"sportin/pkg/models"

	"gorm.io/gorm"
)

type MuscleEntry struct {
	gorm.Model
	Name          string `gorm:"column:name"`
	Description   string `gorm:"column:desciption"`
	MuscleGroupID int    `gorm:"column:muscle_group_id"`
	MuscleGroup   MuscleGroupEntry
}

type MuscleEntryRepository interface {
	Create(muscleEntry *MuscleEntry) (*MuscleEntry, error)
	FindAll() ([]*MuscleEntry, error)
	FindById(id int) (*MuscleEntry, error)
	Update(muscleEntry *MuscleEntry) (*MuscleEntry, error)
	Delete(id int) error
	ToModel(muscleEntry *MuscleEntry) *models.MuscleResponse
	ToModelList(muscleEntries []*MuscleEntry) []*models.MuscleResponse
}

type muscleEntryRepository struct {
	db *gorm.DB
}

func NewMuscleEntryRepository(db *gorm.DB) MuscleEntryRepository {
	return &muscleEntryRepository{db: db}
}

func (r *muscleEntryRepository) Create(muscleEntry *MuscleEntry) (*MuscleEntry, error) {
	if err := r.db.Create(muscleEntry).Error; err != nil {
		return nil, err
	}
	return muscleEntry, nil
}

func (r *muscleEntryRepository) FindAll() ([]*MuscleEntry, error) {
	var muscleEntries []*MuscleEntry
	if err := r.db.Find(&muscleEntries).Error; err != nil {
		return nil, err
	}
	return muscleEntries, nil
}

func (r *muscleEntryRepository) FindById(id int) (*MuscleEntry, error) {
	var muscleEntry MuscleEntry
	if err := r.db.First(&muscleEntry, id).Error; err != nil {
		return nil, err
	}
	return &muscleEntry, nil
}

func (r *muscleEntryRepository) Update(muscleEntry *MuscleEntry) (*MuscleEntry, error) {
	if err := r.db.Save(muscleEntry).Error; err != nil {
		return nil, err
	}
	return muscleEntry, nil
}

func (r *muscleEntryRepository) Delete(id int) error {
	if err := r.db.Delete(&MuscleEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *muscleEntryRepository) ToModel(muscleEntry *MuscleEntry) *models.MuscleResponse {
	return &models.MuscleResponse{
		ID:            int(muscleEntry.ID),
		Name:          muscleEntry.Name,
		Description:   muscleEntry.Description,
		MuscleGroupID: muscleEntry.MuscleGroupID,
	}
}

func (r *muscleEntryRepository) ToModelList(muscleEntries []*MuscleEntry) []*models.MuscleResponse {
	var responses []*models.MuscleResponse
	for _, entry := range muscleEntries {
		responses = append(responses, r.ToModel(entry))
	}
	return responses
}
