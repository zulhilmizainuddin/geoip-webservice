package controllers

import (
	"geoip-webservice/models"

	"github.com/gin-gonic/gin"
)

func IP2LocationController(context *gin.Context) {
	ipaddress := context.Query("ipaddress")

	geolocation := models.IP2LocationQuery(ipaddress)

	context.JSON(200, geolocation)
}
