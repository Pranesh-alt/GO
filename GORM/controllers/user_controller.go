package controllers

import (
	"GORM/models"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	email := r.URL.Query().Get("email")

	var users []models.User
	query := db.Model(&models.User{})
	if email != "" {
		query = query.Where("email = ?", email)
	}

	if err := query.Find(&users).Error; err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"users": users})
}

// Create a user
func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&input).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, input)
}

// Get a user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"user": user})
}

// Update a user
func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	if err := db.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"user": user})
}

// Delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}
