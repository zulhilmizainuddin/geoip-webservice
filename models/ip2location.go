package models

import (
	"path/filepath"

	"github.com/ip2location/ip2location-go"
	"github.com/spf13/viper"
)

type IP2Location struct{}

func (this *IP2Location) Query(ipaddress string) Geolocation {
	path, _ := filepath.Abs(viper.GetString("databases.ip2location"))

	ip2location.Open(path)
	defer ip2location.Close()

	results := ip2location.Get_all(ipaddress)
	geolocation := Geolocation{
		IPAddress: ipaddress,
		City:      results.City,
		Country:   results.Country_long,
		Latitude:  float64(results.Latitude),
		Longitude: float64(results.Longitude),
	}

	return geolocation
}
