package middlewares

import (
	"geoip-webservice/models"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStringValidationMiddleware(c *gin.Context) {
	var ipaddress string

	if ipaddress = c.Query("ipaddress"); ipaddress == "" {
		err := models.Error{
			ErrorMessage: "IP address not provided",
		}

		c.JSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	if net.ParseIP(ipaddress) == nil {
		err := models.Error{
			ErrorMessage: "Invalid IP address provided",
		}

		c.JSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	c.Next()
}
