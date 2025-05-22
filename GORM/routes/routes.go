package routes

import (
	"GORM/controllers"
	"GORM/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	// Public Routes
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.PostLogin(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUsers(w, r, db)
	}).Methods("GET")

	// Protected Routes (Require AuthMiddleware)
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware())

	protected.HandleFunc("/me", controllers.MeHandler).Methods("GET")

	protected.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetProtectedUsers(w, r, db)
	}).Methods("GET")

	protected.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetProtectedUsersByID(w, r, db)
	}).Methods("GET")

	protected.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, db)
	}).Methods("PUT")

	protected.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, db)
	}).Methods("DELETE")

	protected.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Logged out successfully"}`))
	}).Methods("POST")
}
