package routes

import (
	controllers "tempt-go-rest/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userHandler := controllers.NewUserController()
	roleHandler := controllers.NewRoleController()

	// user
	router.POST("/user/register", userHandler.CreateUser)

	// role
	router.POST("/user/role", roleHandler.CreateRole)
	router.GET("/user/role/", roleHandler.GetAllRoles)
	router.GET("/user/role/:id", roleHandler.GetRoleByID)
	router.PUT("/user/role/:id",roleHandler.UpdateRole)
	router.DELETE("/user/role/:id", roleHandler.DeleteRole)
}