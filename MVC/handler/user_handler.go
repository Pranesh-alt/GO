package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yourusername/simple-api/model"
	"github.com/yourusername/simple-api/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *service.UserService
}

// writeJSON is a helper function to write JSON responses.
func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	fmt.Println(("Response "))
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars((r))

	id, err := strconv.Atoi((vars["id"]))
	if err != nil {
		http.Error((w), "Invalid user ID", http.StatusBadRequest)
	}
	user, ok := h.UserService.GetUserByID((id))
	if !ok {
		http.Error((w), "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set(("Content-Type"), "application/json")
	json.NewEncoder((w)).Encode(user)
	fmt.Println("User ID:", id)
	fmt.Println("User found", user)

}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	created := h.UserService.AddUser(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// ... previous imports
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser model.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if updatedUser.Name == "" || updatedUser.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	user, ok := h.UserService.UpdateUser(id, updatedUser)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ok := h.UserService.DeleteUser(id)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
