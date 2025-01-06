package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type MuscleGroupEntry struct {
	gorm.Model
	Name        string         `gorm:"column:name"`
	BodyPart    string         `gorm:"column:body_part"`
	Description string         `gorm:"column:desciption"`
	Level       string         `gorm:"column:level"`
	Muscles     []*MuscleEntry `gorm:"foreignKey:MuscleGroupID"`
}

type MuscleGroupEntryRepository interface {
	Create(muscleGroupEntry *MuscleGroupEntry) (*MuscleGroupEntry, error)
	FindAll() ([]*MuscleGroupEntry, error)
	FindById(id int) (*MuscleGroupEntry, error)
	Update(muscleGroupEntry *MuscleGroupEntry) (*MuscleGroupEntry, error)
	Delete(id int) error
	ToModel(muscleGroupEntry *MuscleGroupEntry) *model.MuscleGroupResponse
	ToModelList(muscleGroupEntries []*MuscleGroupEntry) []*model.MuscleGroupResponse
}

type muscleGroupEntryRepository struct {
	db *gorm.DB
}

func NewMuscleGroupEntryRepository(db *gorm.DB) MuscleGroupEntryRepository {
	return &muscleGroupEntryRepository{db: db}
}

func (r *muscleGroupEntryRepository) Create(muscleGroupEntry *MuscleGroupEntry) (*MuscleGroupEntry, error) {
	if err := r.db.Create(muscleGroupEntry).Error; err != nil {
		return nil, err
	}
	return muscleGroupEntry, nil
}

func (r *muscleGroupEntryRepository) FindAll() ([]*MuscleGroupEntry, error) {
	var entries []*MuscleGroupEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *muscleGroupEntryRepository) FindById(id int) (*MuscleGroupEntry, error) {
	var entry MuscleGroupEntry
	if err := r.db.First(&entry, id).Error; err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *muscleGroupEntryRepository) Update(muscleGroupEntry *MuscleGroupEntry) (*MuscleGroupEntry, error) {
	if err := r.db.Save(muscleGroupEntry).Error; err != nil {
		return nil, err
	}
	return muscleGroupEntry, nil
}

func (r *muscleGroupEntryRepository) Delete(id int) error {
	if err := r.db.Delete(&MuscleGroupEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *muscleGroupEntryRepository) ToModel(muscleGroupEntry *MuscleGroupEntry) *model.MuscleGroupResponse {
	return &model.MuscleGroupResponse{
		ID:          int(muscleGroupEntry.ID),
		Name:        muscleGroupEntry.Name,
		BodyPart:    muscleGroupEntry.BodyPart,
		Description: muscleGroupEntry.Description,
		Level:       muscleGroupEntry.Level,
	}
}

func (r *muscleGroupEntryRepository) ToModelList(muscleGroupEntries []*MuscleGroupEntry) []*model.MuscleGroupResponse {
	var responses []*model.MuscleGroupResponse
	for _, entry := range muscleGroupEntries {
		responses = append(responses, r.ToModel(entry))
	}
	if len(responses) == 0 {
		return []*model.MuscleGroupResponse{}
	}
	return responses
}
