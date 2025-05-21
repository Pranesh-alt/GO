package routes

import (
	"GORM/controllers"
	"GORM/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	// Public
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserByID(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, db)
	}).Methods("DELETE")

	// Protected
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware())

	protected.HandleFunc("/me", controllers.MeHandler).Methods("GET")
}
