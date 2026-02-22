package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Event-driven architecture for modern applications
func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "cloud-fast-connector-ml-0iq",
			"status": "running",
		})
	})
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	
	r.Run(":8080")
}
