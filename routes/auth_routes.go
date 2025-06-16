package routes

import (
	"example.com/goRestAPI/controllers"
	"example.com/goRestAPI/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	protected := router.Group("/api")
	protected.Use(middleware.Authenticate)
	{
		protected.GET("/protected", controllers.ProtectedEndpoint)
		protected.GET("/users", controllers.GetAllUsers)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
