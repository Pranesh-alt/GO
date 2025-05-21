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
	r.GET("/users/:id", func(c *gin.Context) {
		controllers.GetUserByID(c, db)
	})
	r.PUT("/users/:id", func(c *gin.Context) {
		controllers.UpdateUser(c, db)
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		controllers.DeleteUser(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	// Protected routes group
	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		email, _ := c.Get("user_email")
		c.JSON(http.StatusOK, gin.H{"email": email})
	})

	// Role-based protected route example (admin only)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	admin.GET("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	})
}
