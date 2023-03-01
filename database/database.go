package database

import(
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
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
	
	db.AutoMigrate(&Models.card{})


	db = DbInstance{DB: database}
}
