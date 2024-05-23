package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seniorLikeToCode/url-shortener/shortener"
	"github.com/seniorLikeToCode/url-shortener/store"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserID)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserID)

	host := "http://localhost:5000/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShorterUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, _ := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(392, initialUrl)
}
