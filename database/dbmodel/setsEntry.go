package dbmodel

import (
	"sportin/pkg/model"

	"gorm.io/gorm"
)

type SetsEntry struct {
	gorm.Model
	RPE               int    `gorm:"column:rpe"`
	RIR               int    `gorm:"column:rir"`
	Weight            int    `gorm:"column:weight"`
	Work              string `gorm:"column:work"`
	WorkType          string `gorm:"column:work_type"`
	ResistanceBand    string `gorm:"column:resistance_band"`
	RestTime          int    `gorm:"column:rest_time"`
	Intensification   IntensificationEntry
	IntensificationID int `gorm:"column:intensification_id"`
	ProgramExercice   ProgramExerciseEntry
	ProgramExerciceID int `gorm:"column:program_exercice_id"`
}

type SetsEntryRepository interface {
	Create(setsEntry *SetsEntry) (*SetsEntry, error)
	FindAll() ([]*SetsEntry, error)
	FindById(id int) (*SetsEntry, error)
	Update(setsEntry *SetsEntry) (*SetsEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *SetsEntry) *model.SetsReponse
}

type setsEntryRepository struct {
	db *gorm.DB
}

func NewSetsEntryRepository(db *gorm.DB) SetsEntryRepository {
	return &setsEntryRepository{db: db}
}

func (r *setsEntryRepository) Create(setsEntry *SetsEntry) (*SetsEntry, error) {
	if err := r.db.Create(setsEntry).Error; err != nil {
		return nil, err
	}
	return setsEntry, nil
}

func (r *setsEntryRepository) FindAll() ([]*SetsEntry, error) {
	var setsEntries []*SetsEntry
	if err := r.db.Find(&setsEntries).Error; err != nil {
		return nil, err
	}
	return setsEntries, nil
}

func (r *setsEntryRepository) FindById(id int) (*SetsEntry, error) {
	var setsEntry *SetsEntry
	if err := r.db.First(&setsEntry, id).Error; err != nil {
		return nil, err
	}
	return setsEntry, nil
}

func (r *setsEntryRepository) Update(setsEntry *SetsEntry) (*SetsEntry, error) {
	if err := r.db.Where("id = ?", setsEntry.ID).Updates(setsEntry).Error; err != nil {
		return nil, err
	}
	return setsEntry, nil
}

func (r *setsEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&SetsEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *setsEntryRepository) ToModel(entry *SetsEntry) *model.SetsReponse {
	return &model.SetsReponse{
		ID:                int(entry.ID),
		RPE:               entry.RPE,
		RIR:               entry.RIR,
		Weight:            entry.Weight,
		Work:              entry.Work,
		WorkType:          entry.WorkType,
		ResistanceBand:    entry.ResistanceBand,
		RestTime:          entry.RestTime,
		IntensificationID: entry.IntensificationID,
		ProgramExerciseID: entry.ProgramExerciceID,
	}
}
