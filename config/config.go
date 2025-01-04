package config

import (
	"sportin/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(mysql.Open("user:user@tcp(localhost:3306)/sportin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	return &config, nil
}
