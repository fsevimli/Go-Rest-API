# Go REST API - Kullanıcı ve Ürün Yönetimi

Bu proje, **Go**, **Gin Web Framework**, **GORM ORM**, **JWT authentication** ve **SQLite** kullanılarak geliştirilmiş basit bir REST API uygulamasıdır.  
Kullanıcı kayıt/giriş işlemleri, korumalı endpoint'ler, ürün ekleme/güncelleme/silme ve herkese açık ürün listeleme işlemlerini içerir.

---
## Özellikler

- Kullanıcı kayıt ve giriş (JWT ile kimlik doğrulama)
- Token ile korumalı alanlara erişim
- Ürün ekleme, güncelleme, silme (sadece giriş yapan kullanıcılar)
- Ürün listeleme (herkese açık)
- ID ile ürün sorgulama
- Modüler proje yapısı (`models`, `controllers`, `routes`, `middleware`, `utils`)

---
## Kullanılan Teknolojiler

- [Go](https://golang.org/) v1.21+
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) (ORM)
- [JWT](https://github.com/golang-jwt/jwt)
- [SQLite](https://github.com/glebarez/sqlite) (CGO’suz sürüm)

---
## Proje Yapısı
goRestAPI/
│
├── config/ # Veritabanı bağlantısı
├── controllers/ # İş mantıkları (auth, product)
├── middleware/ # JWT doğrulama
├── models/ # Veritabanı modelleri ve işlemleri
├── routes/ # Route tanımlamaları
├── utils/ # Yardımcı fonksiyonlar (Hashing, JWT)
├── main.go # Uygulama giriş noktası
└── go.mod

---
## Kurulum

--```bash
# 1. Bu repoyu klonla
git clone https://github.com/kullaniciadi/goRestAPI.git
cd goRestAPI

# 2. Gerekli modülleri indir
go mod tidy

# 3. Uygulamayı başlat
go run main.go

---
## Örnek İstekler

### Register
POST /auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}

### Login
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}

### Ürün Ekle
POST /api/products/
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Laptop",
  "description": "16GB RAM, 512GB SSD",
  "price": 1500
}

### Ürün Listeleme
GET /api/products
Content-Type: application/json

### Ürün Güncelleme
PUT http://localhost:8080/api/products/4
Content-Type: application/json
Authorization: Bearer <token>

{
  "name": "Kablosuz Mouse",
  "description": "Yüksek DPI",
  "price": 380
}