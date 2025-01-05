package config

import (
	"sportin/database"
	"sportin/database/dbmodel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB             *gorm.DB
	UserRepository dbmodel.UserRepository
	StatsRepository dbmodel.StatsRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(mysql.Open("elemee:elemee123@tcp(localhost:3306)/sportin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.StatsRepository = dbmodel.NewStatsRepository(databaseSession)

	return &config, nil
}
