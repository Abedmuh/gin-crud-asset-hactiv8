package database

import (
	"fmt"
	"tempt-go-rest/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Koneksi() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/go_crud"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Role{})
	if err != nil {
		return nil, err
	}
	DB = db
	fmt.Println("connecting to database")

	return db, nil
}
