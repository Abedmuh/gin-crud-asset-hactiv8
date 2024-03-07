// controllers/Role_controller.go

package controllers

import (
	"fmt"
	"net/http"
	"tempt-go-rest/database"
	"tempt-go-rest/models"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func NewRoleController() *RoleController {
	return &RoleController{}
}

// CreateRole digunakan untuk membuat pengguna baru
func (rc *RoleController) CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	db := database.DB
	result := db.Create(&role)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Berhasil tambah data",
		"role": role})
}

func (rc *RoleController) GetAllRoles(c *gin.Context) {
	db := database.DB

	var roles []models.Role
	result := db.Find(&roles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}


func (rc *RoleController) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	db := database.DB

	var role models.Role
	result := db.First(&role, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (rc *RoleController) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	db := database.DB

	var role models.Role
	result := db.First(&role, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	result = db.Save(&role)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Role with ID %s has been updated", id),
	})
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	db := database.DB
	fmt.Println("mencoba")
	var Role models.Role
	result := db.First(&Role, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	result = db.Delete(&Role)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Role with ID %s has been deleted", id),
	})
}