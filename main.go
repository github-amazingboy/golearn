package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	myGroup := r.Group("/v1")

	myGroup.GET("/ll", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ll",
		})
	})

	r.Run(":8050")
}
