package database

import (
	"tempt-go-rest/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Koneksi() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/go_crud"
	// Ganti user, password, nama database, dan konfigurasi sesuai dengan informasi koneksi MySQL Anda

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Role{})
	if err != nil {
		return nil, err
	}
	DB = db

	return db, nil
}
