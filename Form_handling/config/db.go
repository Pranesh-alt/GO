package config

import (
	"Form_handling/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:62145090@tcp(127.0.0.1:3306)/Form_handling?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("DB connected")

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
}
