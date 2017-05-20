package controllers

import (
	"geoip-webservice/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MaxMindController(c *gin.Context) {
	ipaddress := c.Query(viper.GetString("queryStrings.ipaddress"))

	maxmind := models.MaxMind{}
	geolocation := maxmind.Query(ipaddress)

	c.JSON(http.StatusOK, geolocation)
}
