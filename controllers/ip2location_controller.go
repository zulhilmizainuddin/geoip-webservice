package controllers

import (
	"github.com/zulhilmizainuddin/geoip-webservice/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IP2LocationController(c *gin.Context) {
	ipaddress := c.Query(viper.GetString("querystrings.ipaddress"))

	ip2location := models.IP2Location{}
	geolocation := ip2location.Query(ipaddress)

	c.JSON(http.StatusOK, geolocation)
}
