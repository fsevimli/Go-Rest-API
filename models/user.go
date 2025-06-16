package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/goRestAPI/config"
	"example.com/goRestAPI/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Kullanıcıyı kaydederken hashle
func (u *User) Save() error {

	// Önce email var mı kontrol et
	var count int64
	config.DB.Table("users").Where("email = ?", u.Email).Count(&count)
	if count > 0 {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	now := time.Now()
	result := config.DB.Exec(query, u.Name, u.Email, u.Password, now, now)

	return result.Error
}

func (u *User) ValidateCredentials() error {
	// E-posta ile kullanıcıyı veritabanından sorgula
	fmt.Println("Check user with email:", u.Email)
	query := "SELECT id, password FROM users WHERE email = ?"
	row := config.DB.Raw(query, u.Email).Row()

	var hashedPassword string
	var userID uint

	// Sonuçları değişkenlere aktar
	err := row.Scan(&userID, &hashedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	fmt.Println("Hashed password from DB:", hashedPassword)
	fmt.Println("Plain password from user:", u.Password)

	// Girilen şifreyle hashlenmiş şifreyi karşılaştır
	passwordIsValid := utils.CheckPasswordHash(u.Password, hashedPassword)
	if !passwordIsValid {
		fmt.Println("Password check failed")
		return errors.New("credentials invalid")
	}

	fmt.Println("Login başarılı")
	// Giriş başarılıysa User modelini doldur
	u.ID = userID
	u.Password = hashedPassword

	return nil
}
