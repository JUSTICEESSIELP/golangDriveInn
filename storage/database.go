package storage

import (
	"driveinn_server/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() *gorm.DB {

	envError := godotenv.Load()
	if envError != nil {
		panic("Error laoding .env file")
	}
	dbUrl := os.Getenv("DB_CONNECTION_STRING")
	db, dbError := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if dbError != nil {
		log.Panic("error connecting to db")

	}

	DB = db
	return db

	// the main thing is to link the env db url and connect with iris syntax
}

func performMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func InitializeDB() *gorm.DB {

	db := connectToDB()
	performMigrations(db)
	return db
}
