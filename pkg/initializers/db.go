package initializers

import (
	"fmt"
	"go-mvc/pkg/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Task{})
}
