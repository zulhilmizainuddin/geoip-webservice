package models

type Geolocation struct {
	IPAddress string  `json:"ipaddress"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
