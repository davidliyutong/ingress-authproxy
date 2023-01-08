package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "love",
		"user2": "god",
		"user3": "sex",
	}))

	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})

	_ = r.Run() // 监听服务在 0.0.0.0:8080
}
