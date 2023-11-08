package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isaacodekill/rate-limiterr/core"
)

type LimitedEndpointController struct {
	rateLimiter core.RateLimiter
}

func NewLimitedEndpointController(rateLimiter core.RateLimiter) LimitedEndpointController {
	return LimitedEndpointController{rateLimiter: rateLimiter}
}

func (LEC LimitedEndpointController) testRequest(c *gin.Context) bool {
	ip := c.ClientIP()
	if !LEC.rateLimiter.AllowRequest(ip) {
		c.JSON(429, gin.H{
			"message": "too many requests",
		})
		return false
		// close the response writer
	}
	return true
}

func (LEC LimitedEndpointController) LimitedEndpoint(c *gin.Context) {
	allowReq := LEC.testRequest(c)
	fmt.Println(c.Request)

	if allowReq {
		c.JSON(200, gin.H{
			"message": "limited endpoint",
		})
	}
}

func UnlimitedEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "unlimited endpoint",
	})
}
