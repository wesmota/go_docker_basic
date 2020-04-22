package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	public := r.Group("/public")

	public.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})

	})
	r.Run("localhost:8080")
}
