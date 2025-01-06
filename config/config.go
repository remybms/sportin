package config

import (
	"sportin/database"
	"sportin/database/dbmodel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB                             *gorm.DB
	MuscleGroupEntryRepository     dbmodel.MuscleGroupEntryRepository
	UserRepository                 dbmodel.UserRepository
	UserStatsRepository            dbmodel.UserStatsRepository
	CategoryEntryRepository        dbmodel.CategoryEntryRepository
	ProgramEntryRepository         dbmodel.ProgramEntryRepository
	MuscleEntryRepository          dbmodel.MuscleEntryRepository
	IntensificationEntryRepository dbmodel.IntensificationEntryRepository
	ExerciseEntryRepository    dbmodel.ExerciseEntryRepository

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(mysql.Open("user:user@tcp(localhost:3306)/sportin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)
	config.CategoryEntryRepository = dbmodel.NewCategoryRepository(databaseSession)
	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.UserStatsRepository = dbmodel.NewUserStatsRepository(databaseSession)
	config.MuscleGroupEntryRepository = dbmodel.NewMuscleGroupEntryRepository(databaseSession)
	config.MuscleEntryRepository = dbmodel.NewMuscleEntryRepository(databaseSession)
	config.ProgramEntryRepository = dbmodel.NewProgramEntryRepository(databaseSession)
	config.MuscleEntryRepository = dbmodel.NewMuscleEntryRepository(databaseSession)
	config.ExerciseEntryRepository = dbmodel.NewExerciseEntryRepository(databaseSession)
	config.IntensificationEntryRepository = dbmodel.NewIntensificationEntryRepository(databaseSession)

	return &config, nil
}
