package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Nissekissen/GO-testing/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var DB DBinstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Stockholm",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Token{})

	DB = DBinstance{
		DB: db,
	}

}
