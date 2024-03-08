package main

import (
	"fmt"
	"net/http"
	"tempt-go-rest/database"
	"tempt-go-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Koneksi()
	if err!= nil {
    fmt.Println(err)
		return
  }	
	
	r := gin.Default()
	routes.UserRoutes(r,db)
	
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	r.Run(":3000") 
}