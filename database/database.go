package database

import (
	"github.com/AlexandreHardyy/goProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=example dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Payment{}, &models.Product{})

	if err != nil {
		return
	}

	DB = database
}
