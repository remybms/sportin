package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type ProgramEntry struct {
	gorm.Model
	Name        string        `gorm:"column:name"`
	Description string        `gorm:"column:description"`
	UserID      int           `gorm:"column:user_id"`
	User        UserEntry     `gorm:"foreignKey:UserID"`
	CategoryID  int           `gorm:"column:category_id"`
	Category    CategoryEntry `gorm:"foreignKey:CategoryID"`
}

type ProgramEntryRepository interface {
	Create(programEntry *ProgramEntry) (*ProgramEntry, error)
	FindAll() ([]*ProgramEntry, error)
	FindByID(id int) (*ProgramEntry, error)
	Update(programEntry *ProgramEntry) (*ProgramEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *ProgramEntry) *model.ProgramResponse
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
	if err := r.db.First(&programEntry, id).Error; err != nil {
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

func (r *programEntryRepository) ToModel(entry *ProgramEntry) *model.ProgramResponse {
	return &model.ProgramResponse{
		ID:          int(entry.ID),
		Name:        entry.Name,
		Description: entry.Description,
	}
}
