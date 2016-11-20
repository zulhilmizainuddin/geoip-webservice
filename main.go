package main

import (
	"geoip-webservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ip2location", controllers.IP2LocationController)

	r.Run(":3000")
}
