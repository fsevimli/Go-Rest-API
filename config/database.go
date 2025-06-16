package config

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}

	DB = database
	log.Println("Veritabanına başarıyla bağlanıldı.")
}
