package main

import (
	"geoip-webservice/controllers"
	"geoip-webservice/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.QueryStringValidationMiddleware)

	r.GET("/ip2location", controllers.IP2LocationController)

	r.Run(":3000")
}
