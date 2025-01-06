package database

import (
	"log"
	"sportin/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.MuscleGroupEntry{},
	)
	if err != nil {
		log.Panicln("Database migration failed:", err)
	}
	log.Println("Database migrated successfully")
}
