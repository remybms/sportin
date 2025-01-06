package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type CategoryEntry struct {
	gorm.Model
	Name        string          `gorm:"column:name"`
	Description string          `gorm:"column:description"`
	Programs    []*ProgramEntry `gorm:"foreignKey:CategoryID"`
}

type CategoryEntryRepository interface {
	Create(categoryEntry *CategoryEntry) (*CategoryEntry, error)
	FindAll() ([]*CategoryEntry, error)
	FindById(id int) (*CategoryEntry, error)
	Update(categoryEntry *CategoryEntry) (*CategoryEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *CategoryEntry) *model.CategoryResponse
}

type categoryEntryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryEntryRepository {
	return &categoryEntryRepository{db: db}
}

func (r *categoryEntryRepository) Create(categoryEntry *CategoryEntry) (*CategoryEntry, error) {
	if err := r.db.Create(categoryEntry).Error; err != nil {
		return nil, err
	}
	return categoryEntry, nil
}

func (r *categoryEntryRepository) FindAll() ([]*CategoryEntry, error) {
	var categoryEntries []*CategoryEntry
	if err := r.db.Find(&categoryEntries).Error; err != nil {
		return nil, err
	}
	return categoryEntries, nil
}

func (r *categoryEntryRepository) FindById(id int) (*CategoryEntry, error) {
	var categoryEntry *CategoryEntry
	if err := r.db.First(&categoryEntry, id).Error; err != nil {
		return nil, err
	}
	return categoryEntry, nil
}

func (r *categoryEntryRepository) Update(categoryEntry *CategoryEntry) (*CategoryEntry, error) {
	if err := r.db.Where("id = ?", categoryEntry.ID).Updates(categoryEntry).Error; err != nil {
		return nil, err
	}
	return categoryEntry, nil
}

func (r *categoryEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&CategoryEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *categoryEntryRepository) ToModel(entry *CategoryEntry) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:          int(entry.ID),
		Name:        entry.Name,
		Description: entry.Description,
	}
}
