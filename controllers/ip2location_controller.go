package controllers

import (
	"geoip-webservice/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IP2LocationController(c *gin.Context) {
	ipaddress := c.Query(viper.GetString("querystrings.ipaddress"))

	geolocation := models.IP2LocationQuery(ipaddress)

	c.JSON(http.StatusOK, geolocation)
}
