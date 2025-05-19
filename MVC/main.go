package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yourusername/simple-api/handler"
	"github.com/yourusername/simple-api/service"
	"net/http"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	// Use gorilla/mux router
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/health", userHandler.HealthCheck).Methods("GET")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/search", userHandler.SearchUsers).Methods("GET")
	r.HandleFunc("/users/email/{email}", userHandler.GetUserByEmail).Methods("GET")
	r.HandleFunc("/users/stats", userHandler.GetUserStats).Methods("GET")
	r.HandleFunc("/users", userHandler.DeleteUserByEmail).Methods("DELETE")

	// Enable CORS
	corsHandler := cors.Default().Handler(r)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler)
}
