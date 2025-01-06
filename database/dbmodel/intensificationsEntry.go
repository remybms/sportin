package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type IntensificationEntry struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

type IntensificationEntryRepository interface {
	Create(intensificationEntry *IntensificationEntry) (*IntensificationEntry, error)
	FindAll() ([]*IntensificationEntry, error)
	FindById(id int) (*IntensificationEntry, error)
	Update(intensificationEntry *IntensificationEntry) (*IntensificationEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *IntensificationEntry) *model.IntensificationResponse
}

type intensificationEntryRepository struct {
	db *gorm.DB
}

func NewIntensificationEntryRepository(db *gorm.DB) IntensificationEntryRepository {
	return &intensificationEntryRepository{db: db}
}

func (r *intensificationEntryRepository) Create(intensificationEntry *IntensificationEntry) (*IntensificationEntry, error) {
	if err := r.db.Create(intensificationEntry).Error; err != nil {
		return nil, err
	}
	return intensificationEntry, nil
}

func (r *intensificationEntryRepository) FindAll() ([]*IntensificationEntry, error) {
	var intensificationEntries []*IntensificationEntry
	if err := r.db.Find(&intensificationEntries).Error; err != nil {
		return nil, err
	}
	return intensificationEntries, nil
}

func (r *intensificationEntryRepository) FindById(id int) (*IntensificationEntry, error) {
	var intensificationEntry *IntensificationEntry
	if err := r.db.First(&intensificationEntry, id).Error; err != nil {
		return nil, err
	}
	return intensificationEntry, nil
}

func (r *intensificationEntryRepository) Update(intensificationEntry *IntensificationEntry) (*IntensificationEntry, error) {
	if err := r.db.Where("id = ?", intensificationEntry.ID).Updates(intensificationEntry).Error; err != nil {
		return nil, err
	}
	return intensificationEntry, nil
}

func (r *intensificationEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&IntensificationEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *intensificationEntryRepository) ToModel(entry *IntensificationEntry) *model.IntensificationResponse {
	return &model.IntensificationResponse{
		ID:          int(entry.ID),
		Name:        entry.Name,
		Description: entry.Description,
	}
}
