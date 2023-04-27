package middlewares

import (
	"net"
	"net/http"

	"github.com/zulhilmizainuddin/geoip-webservice/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func QueryStringValidationMiddleware(c *gin.Context) {
	var ipaddress string

	if ipaddress = c.Query(viper.GetString("querystrings.ipaddress")); ipaddress == "" {
		err := models.Error{
			ErrorMessage: viper.GetString("errors.ip_address_not_provided"),
		}

		c.JSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	if net.ParseIP(ipaddress) == nil {
		err := models.Error{
			ErrorMessage: viper.GetString("errors.invalid_ip_address_provided"),
		}

		c.JSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	c.Next()
}
