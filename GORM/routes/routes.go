package routes

import (
	"GORM/controllers"
	"GORM/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	// Public routes
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserByID(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, db)
	}).Methods("DELETE")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(w, r, db)
	}).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware())

	protected.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value("user_email").(string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"email":"` + email + `"}`))
	}).Methods("GET")

	// Admin routes
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware("admin"))

	admin.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Welcome Admin"}`))
	}).Methods("GET")
}
