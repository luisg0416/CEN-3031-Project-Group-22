package storage

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Congig struct(
	Host	 string
	Port	 string
	Password string
	User	 string
	DBName	 string
	SSLMode	 string
)

func NewConnection(config *Config)(*gorm.DB, error){
	dsn : fmt.Sprintf(
		"host=%s port=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gormOpen(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return db, err
	}
	return db, nil
}