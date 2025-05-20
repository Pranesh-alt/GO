package routes

import (
	"GORM/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})
}
