package main

import (
	"example.com/goRestAPI/config"
	"example.com/goRestAPI/models"
	"example.com/goRestAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Product{})

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API çalışıyor."})
	})

	routes.RegisterAuthRoutes(server)
	routes.ProductRoutes(server)

	server.Run(":8080") //localhost:8080
}
