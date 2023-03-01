package database

import(
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type DbInstance struct {
	Db *gorm.Db
}

var db DbInstance

func ConnectDb() {
	database, err := gorm.Open(sqlite.Open("api.db"), &gorm.config{})

	if err != nil {
		log.Fatal("Database connection Failes!", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to Database")
	database.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Migrations
	db.AutoMigrate(&Models.card{})

	db = DbInstance{Db: database}
}
