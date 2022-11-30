package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/untletch/books-crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(uri string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("error in connecting to db: %v", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})

	return db
}
