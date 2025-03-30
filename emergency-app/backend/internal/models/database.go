package models

import (
	"gorm.io/gorm"
)

// DB is a shared database instance
var DB *gorm.DB

// SetDB sets the database instance for models
func SetDB(db *gorm.DB) {
	DB = db
}