package connector

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"example.com/RestAPIgo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var books []models.Book

	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&books)
	if err != nil {
		return
	}

	DB = database
}
