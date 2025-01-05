package database

import (
	"log"
	"sportin/database/dbmodel"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&dbmodel.MuscleGroupEntry{})
	log.Println("Database migrated successfully")
}
