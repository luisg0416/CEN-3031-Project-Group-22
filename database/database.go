package database

import (
	"log"
	"os"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var db DbInstance

func ConnectDb() {

	database, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection Failed", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to Database")
	database.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	
	// Migrations
	db.DB.AutoMigrate(&Models.Card{})

	db = DbInstance{DB: database}
}
