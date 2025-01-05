package dbmodel

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	UserId              int `json:"user_id"`
	Weight              int `json:"user_weight"`
	Height              int `json:"user_height"`
	Age                 int `json:"user_age"`
	ActivityCoefficient int `json:"user_activity"`
	CaloriesGoal        int `json:"user_calories_goal"`
	ProteinRatio        int `json:"user_protein_ratio"`
}

type StatsRepository interface {
	Create(newStats *Stats) (*Stats, error)
	FindAll() ([]*Stats, error)
	FindById(statsId string) ([]*Stats, error)
	Delete(statsToDelete *Stats) error
	Update(statsToUpdate *Stats, statId string) error
}

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) Create(stats *Stats) (*Stats, error) {
	if err := r.db.Create(stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

func (r *statsRepository) Delete(statsToDelete *Stats) error {
	if err := r.db.Delete(statsToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *statsRepository) FindAll() ([]*Stats, error) {
	var stats []*Stats
	if err := r.db.Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

func (r *statsRepository) FindById(statsId string) ([]*Stats, error) {
	var entry []*Stats
	if err := r.db.Where("id = ?", statsId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *statsRepository) Update(statsToUpdate *Stats, statsId string) error {
	if err := r.db.Where("id = ?", statsId).Updates(statsToUpdate).Error; err != nil {
		return err
	}
	return nil
}
