package main

import (
	"geoip-webservice/controllers"
	"geoip-webservice/middlewares"
	"geoip-webservice/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	if err := utils.ReadConfigFile(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(middlewares.QueryStringValidationMiddleware)

	r.GET(viper.GetString("routes.ip2location"), controllers.IP2LocationController)
	r.GET(viper.GetString("routes.maxmind"), controllers.MaxMindController)

	r.Run(viper.GetString("server.port"))
}
