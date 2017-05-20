package models

import (
	"net"
	"path/filepath"

	"github.com/oschwald/geoip2-golang"
	"github.com/spf13/viper"
)

type MaxMind struct{}

func (this *MaxMind) Query(ipaddress string) Geolocation {
	path, _ := filepath.Abs(viper.GetString("databases.maxmind"))

	db, _ := geoip2.Open(path)
	defer db.Close()

	ip := net.ParseIP(ipaddress)
	results, _ := db.City(ip)
	geolocation := Geolocation{
		IPAddress: ipaddress,
		City:      results.City.Names["en"],
		Country:   results.Country.Names["en"],
		Latitude:  results.Location.Latitude,
		Longitude: results.Location.Longitude,
	}

	return geolocation
}
