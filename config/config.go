package config

import (
	"fmt"
	"log"
	"os"
	"sportin/database"
	"sportin/database/dbmodel"

	"github.com/joho/godotenv"
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

	user := goDotEnvVariable("DBUSER")
	password := goDotEnvVariable("DBPASSWORD")
	port := goDotEnvVariable("DBPORT")
	host := goDotEnvVariable("DBHOST")
	dbName := goDotEnvVariable("DBNAME")

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

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Print(err)
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
