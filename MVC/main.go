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
	// Initialize router
	r := mux.NewRouter()

	// Existing routes...
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.GetUsers(w, r)
		} else if r.Method == http.MethodPost {
			userHandler.CreateUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	handler := cors.Default().Handler(mux)

	fmt.Println("Server started on http//localhost:8081")
	http.ListenAndServe(":8080", handler)
}
