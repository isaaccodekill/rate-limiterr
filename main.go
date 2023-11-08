package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacodekill/rate-limiterr/controller"
	"github.com/isaacodekill/rate-limiterr/core"
)

func main() {

	r := gin.Default()

	// create a new rate limiter
	rateLimiter := core.NewRateLimiter(core.TokenBucketLimiterType)

	// create a new limited endpoint controller
	limitedEndpointController := controller.NewLimitedEndpointController(rateLimiter)

	r.GET("/limited", limitedEndpointController.LimitedEndpoint)
	r.GET("/unlimited", controller.UnlimitedEndpoint)
	r.Run(":8080")
}
