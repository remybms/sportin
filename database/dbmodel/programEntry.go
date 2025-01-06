package dbmodel

import (
	"sportin/pkg/models"

	"gorm.io/gorm"
)

type ProgramEntry struct {
	gorm.Model
	Name        string `gorm:"colmn:name`
	Description string `gorm:"column:description`

	Users      []*User       `gorm:"foreignKey:UserId"`
	Categories []*Categories `gorm:"foreignKey:CategoryId"`
}

type ProgramEntryRepository interface {
	Create(programEntry *ProgramEntry) (*ProgramEntry, error)
	FindAll() ([]*ProgramEntry, error)
	FindByID(id int) (*ProgramEntry, error)
	Update(programEntry *ProgramEntry) (*ProgramEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *ProgramEntry) *models.ProgramResponse
}

type programEntryRepository struct {
	db *gorm.DB
}

func NewProgramEntryRepository(db *gorm.DB) ProgramEntryRepository {
	return &programEntryRepository{db: db}
}

func (r *programEntryRepository) Create(programEntry *ProgramEntry) (*ProgramEntry, error) {
	if err := r.db.Create(programEntry).Error; err != nil {
		return nil, err
	}
	return programEntry, nil
}

func (r *programEntryRepository) FindAll() ([]*ProgramEntry, error) {
	var programEntries []*ProgramEntry
	if err := r.db.Find(&programEntries).Error; err != nil {
		return nil, err
	}
	return programEntries, nil
}

func (r *programEntryRepository) FindByID(id int) (*ProgramEntry, error) {
	var programEntry *ProgramEntry
	if err := r.db.Where("id = ?", id).Find(&programEntry).Error; err != nil {
		return nil, err
	}
	return programEntry, nil
}

func (r *programEntryRepository) Update(programEntry *ProgramEntry) (*ProgramEntry, error) {
	if err := r.db.Where("id = ?", programEntry.ID).Updates(&programEntry).Error; err != nil {
		return nil, err
	}
	return programEntry, nil
}

func (r *programEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&ProgramEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *programEntryRepository) ToModel(entry *ProgramEntry) *models.ProgramResponse {
	return &models.ProgramResponse{
		ID:          int(entry.ID),
		Name:        entry.Name,
		Description: entry.Description,
	}
}
