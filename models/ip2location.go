package models

import (
	"path/filepath"

	"github.com/ip2location/ip2location-go"
)

func IP2LocationQuery(ipaddress string) interface{} {
	path, _ := filepath.Abs("./databases/ip2location/IP2LOCATION-LITE-DB5.IPV6.BIN")

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
