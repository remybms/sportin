package database

import (
	"log"
	"sportin/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dbmodel.UserEntry{},
		&dbmodel.UserStatsEntry{},
		&dbmodel.MuscleGroupEntry{},
		&dbmodel.CategoryEntry{},
		&dbmodel.ProgramEntry{},
	)
	if err != nil {
		log.Panicln("Database migration failed:", err)
	}
	log.Println("Database migrated successfully")
}
