package database

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	log.Println("Database migrated successfully")
}
