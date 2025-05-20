package controllers

import (
	"GORM/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
