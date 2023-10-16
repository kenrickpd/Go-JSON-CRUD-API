package initializers

import (
	"log"

	"github.com/kenrickpd/Go-JSON-CRUD-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
	db.AutoMigrate(&models.User{})
}
