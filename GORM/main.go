package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	_ "GORM/docs" // replace with actual module name
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"GORM/routes"
)

var DB *gorm.DB

func initDB() {
	dsn := "root:62145090@tcp(127.0.0.1:3306)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
}

func main() {
	initDB()

	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	routes.RegisterUserRoutes(r, DB)

	r.Run(":8080")
}
