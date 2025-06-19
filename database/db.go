package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ramnakm/account-service/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=accountdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	database.AutoMigrate(&models.Account{})
	DB = database
}
