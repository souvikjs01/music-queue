package server

import (
	"log"
	"music-queue/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(db string) {
	var err error
	DB, err = gorm.Open(postgres.Open(db), &gorm.Config{})

	if err != nil {
		panic("Failed to connect db")
	}

	log.Println("Database connected successfully")
	DB.AutoMigrate(&models.Song{})
}
