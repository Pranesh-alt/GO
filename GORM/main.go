package main

import (
	"GORM/routes"
	"log"
	"net/http"

	_ "GORM/docs" // Swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	// Swagger docs route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Register application routes
	routes.RegisterUserRoutes(r, DB)

	// Start server
	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
