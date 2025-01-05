package dbmodel

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name        string `json:"categories_name"`
	Description string `json:"categories_description"`
}

type CategoriesRepository interface {
	Create(NewCategory *Categories) (*Categories, error)
	FindAll() ([]*Categories, error)
	FindById(catId string) ([]*Categories, error)
	Delete(catToDelete *Categories) error
	Update(catToUpdate *Categories, catId string) error
}

type categoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {
	return &categoriesRepository{db: db}
}

func (r *categoriesRepository) Create(category *Categories) (*Categories, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoriesRepository) Delete(categoryToDelete *Categories) error {
	if err := r.db.Delete(categoryToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoriesRepository) FindAll() ([]*Categories, error) {
	var categories []*Categories
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoriesRepository) FindById(categoryId string) ([]*Categories, error) {
	var entry []*Categories
	if err := r.db.Where("id = ?", categoryId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *categoriesRepository) Update(categoryToUpdate *Categories, categoryId string) error {
	if err := r.db.Where("id = ?", categoryId).Updates(categoryToUpdate).Error; err != nil {
		return err
	}
	return nil
}
