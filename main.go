package main

import (
	"fmt"
	"net/http"
	"tempt-go-rest/database"
	"tempt-go-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.Koneksi()
	if err!= nil {
    fmt.Println(err)
		return
  }	
	fmt.Println("connecting to database")
	
	r := gin.Default()
	routes.UserRoutes(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}