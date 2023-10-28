package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *grom.DB

func InitDB(dbName string) *grom.DB {
	db, err := grom.Open(sqlite.Open(dbName), &grom.config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	DB = db
	return db
}

func Migrate(db *grom.DB) {
	db.AutoMirate(&Book{})
}
