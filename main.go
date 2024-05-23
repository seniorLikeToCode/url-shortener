package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello Go URL shortener !🚀")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	err := r.Run(":5000")
	if err != nil {
		log.Fatalf("failed to start the web server - Error: %v", err)
	}
}
