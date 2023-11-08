package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacodekill/rate-limiterr/controller"
)

func main() {

	r := gin.Default()
	r.GET("/limited", controller.LimitedEndpoint)
	r.GET("/unlimited", controller.UnlimitedEndpoint)
	r.Run(":8080")
}
