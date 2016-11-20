package controllers

import (
	"geoip-webservice/models"

	"github.com/gin-gonic/gin"
)

func IP2LocationController(c *gin.Context) {
	ipaddress := c.Query("ipaddress")

	geolocation := models.IP2LocationQuery(ipaddress)

	c.JSON(200, geolocation)
}
