package controllers

import (
	"net/http"
	"strconv"

	"example.com/goRestAPI/config"
	"example.com/goRestAPI/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	var product models.Product

	// Json dan veriyi çözümle
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Geçersiz ürün verisi: " + err.Error(),
		})
		return
	}

	// Veritabanına ürünü kaydet
	result := config.DB.Create(&product)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün kaydedilemedi."})
		return
	}

	context.JSON(http.StatusCreated, product)
}

func GetProducts(context *gin.Context) {
	var products []models.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ürünler getirilemedi."})
		return
	}

	context.JSON(http.StatusOK, products)
}

func DeleteProduct(context *gin.Context) {
	id := context.Param("id")

	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün bulunamadı!"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün silinemedi."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Ürün başarıyla silindi!"})
}

func UpdateProduct(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ürün ID'si"})
		return
	}

	var product models.Product

	// Güncellenecek ürünü bul
	if err := config.DB.First(&product, uint(id)).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Ürün bulunamadı"})
		return
	}

	// Gelen JSON verisini oku
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri: " + err.Error()})
		return
	}

	// Ürünü güncelle
	if err := config.DB.Save(&product).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün güncellenemedi"})
		return
	}

	context.JSON(http.StatusOK, product)
}

func GetProductByID(context *gin.Context) {
	id := context.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Ürün bulunamadı."})
		return
	}

	context.JSON(http.StatusOK, product)
}
