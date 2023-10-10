package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	r.Run(":" + port)
}
