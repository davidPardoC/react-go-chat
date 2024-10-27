package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	fmt.Println("Starting - API")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
