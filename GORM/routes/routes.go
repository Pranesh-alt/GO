package routes

import (
	"GORM/controllers"
	"GORM/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	// Public routes
	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})

	// Protected route
	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/me", func(c *gin.Context) {
		email := c.MustGet("user_email").(string)
		c.JSON(http.StatusOK, gin.H{"email": email})
	})
}
