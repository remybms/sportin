package config

import (
	"sportin/database"
	"sportin/database/dbmodel"
	"sportin/helper"

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
	ExerciseEntryRepository        dbmodel.ExerciseEntryRepository
	SetsEntryRepository            dbmodel.SetsEntryRepository
	ProgramExerciseEntryRepository dbmodel.ProgramExerciseEntryRepository
}

func New() (*Config, error) {
	config := Config{}

	user := helper.GoDotEnvVariable("DBUSER")
	password := helper.GoDotEnvVariable("DBPASSWORD")
	port := helper.GoDotEnvVariable("DBPORT")
	host := helper.GoDotEnvVariable("DBHOST")
	dbName := helper.GoDotEnvVariable("DBNAME")

	databaseSession, err := gorm.Open(mysql.Open(user+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
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
	config.SetsEntryRepository = dbmodel.NewSetsEntryRepository(databaseSession)

	config.ProgramExerciseEntryRepository = dbmodel.NewProgramExerciseEntryRepository(databaseSession)
	return &config, nil
}
