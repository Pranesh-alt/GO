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
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
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
