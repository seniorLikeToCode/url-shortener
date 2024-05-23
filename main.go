package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/seniorLikeToCode/url-shortener/handler"
	"github.com/seniorLikeToCode/url-shortener/store"
)

func main() {
	fmt.Printf("Hello Go URL shortener !🚀")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShorterUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":5000")
	if err != nil {
		log.Fatalf("failed to start the web server - Error: %v", err)
	}
}
