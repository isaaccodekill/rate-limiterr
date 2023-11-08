package controller

import "github.com/gin-gonic/gin"

func LimitedEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "limited endpoint",
	})
}

func UnlimitedEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "unlimited endpoint",
	})
}
