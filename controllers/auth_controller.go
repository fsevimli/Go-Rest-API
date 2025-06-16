package controllers

import (
	"net/http"

	"example.com/goRestAPI/models"
	"example.com/goRestAPI/utils"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri: " + err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Kullanıcı kaydedilemedi: " + err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Token üretilemedi."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Kullanıcı başarıyla oluşturuldu.",
		"token":   token,
	})

}

func Login(context *gin.Context) {
	var user models.User

	// Gelen veriyi çöz
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Giriş bilgilerini doğrula
	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	// Token oluştur
	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

func ProtectedEndpoint(context *gin.Context) {
	userId := context.MustGet("userId")
	context.JSON(http.StatusOK, gin.H{
		"message": "Korumalı endpointe hoş geldiniz!",
		"userId":  userId,
	})
}
