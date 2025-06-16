package controllers

import (
	"net/http"

	"example.com/goRestAPI/config"
	"example.com/goRestAPI/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	var users []models.User

	result := config.DB.Select("id", "name", "email", "created_at").Find(&users)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Kullanıcılar getirilemedi.",
		})
		return
	}
	context.JSON(http.StatusOK, users)
}
