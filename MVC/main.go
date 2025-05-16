package main

import (
	"github.com/rs/cors"
	"github.com/yourusername/simple-api/handler"
	"github.com/yourusername/simple-api/service"
	"net/http"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

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
	http.ListenAndServe(":8080", handler)
}
