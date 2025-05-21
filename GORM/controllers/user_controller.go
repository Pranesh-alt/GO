package controllers

import (
	"GORM/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// Get all users
func GetUsers(c *gin.Context, db *gorm.DB) {
	email := c.Query("email")

	var users []models.User

	query := db.Model(&models.User{})

	if email != "" {
		query = query.Where("email = ?", email)
	}

	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Create a user
func CreateUser(c *gin.Context, db *gorm.DB) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// Get a user by ID
func GetUserByID(c *gin.Context, db *gorm.DB) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid user ID param: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		log.Printf("Failed to find user with ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update  a user
func UpdateUser(c *gin.Context, db *gorm.DB) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid user ID param: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		log.Printf("Failed to find user with ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	var input models.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	if err := db.Save(&user).Error; err != nil {
		log.Printf("Failed to update user with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Delete a user
func DeleteUser(c *gin.Context, db *gorm.DB) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid user ID param: %v", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		log.Printf("Failed to find user with ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		log.Printf("Failed to delete user with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
