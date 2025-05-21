package main

import (
	"GORM/routes"
	"log"
	"net/http"

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
		log.Fatal("Failed to connect to database:", err)
	}
}

func main() {
	initDB()

	r := mux.NewRouter()
	routes.RegisterUserRoutes(r, DB)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
