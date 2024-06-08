package internal

import (
	"autoservice-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "autoservice_user:password123@tcp(127.0.0.1:3306)/autoservice?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established successfully")

	// Автоматическая миграция моделей
	err = DB.AutoMigrate(
		&models.Client{},
		&models.Employee{},
		&models.Service{},
		&models.Appointment{},
		&models.News{},
		&models.Contact{},
		&models.Feedback{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	log.Println("Database migrated successfully")
}
