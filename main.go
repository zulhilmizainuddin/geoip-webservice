package main

import (
	"geoip-webservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ip2location", controllers.IP2LocationController)

	router.Run(":3000")
}
