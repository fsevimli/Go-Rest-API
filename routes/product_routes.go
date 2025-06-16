package routes

import (
	"example.com/goRestAPI/controllers"
	"example.com/goRestAPI/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	api := router.Group("/api/products")

	// Herkese açık işlemler
	api.GET("/", controllers.GetProducts)
	api.GET("/:id", controllers.GetProductByID)

	// Giriş yapan kullanıcılar için korumalı işlemler
	api.Use(middleware.Authenticate)
	{
		api.POST("/", controllers.CreateProduct)
		api.PUT("/:id", controllers.UpdateProduct)
		api.DELETE("/:id", controllers.DeleteProduct)
	}
}
