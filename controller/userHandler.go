// controllers/user_controller.go

package controllers

import (
	"net/http"
	"tempt-go-rest/database"
	"tempt-go-rest/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// CreateUser digunakan untuk membuat pengguna baru
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan pengguna ke database
	db := database.DB
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": result,
		"data": user})
}
