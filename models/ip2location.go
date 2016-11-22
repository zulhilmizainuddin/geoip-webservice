package models

import (
	"path/filepath"

	"github.com/ip2location/ip2location-go"
	"github.com/spf13/viper"
)

func IP2LocationQuery(ipaddress string) interface{} {
	path, _ := filepath.Abs(viper.GetString("databases.ip2location"))

	ip2location.Open(path)

	results := ip2location.Get_all(ipaddress)
	geolocation := Geolocation{
		IPAddress: ipaddress,
		City:      results.City,
		Country:   results.Country_long,
		Latitude:  results.Latitude,
		Longitude: results.Longitude,
	}

	ip2location.Close()

	return geolocation
}
