package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router = gin.Default()
)

func main() {
	//router.Run("8080")
	//router.GET("/ping", ping.Ping)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
