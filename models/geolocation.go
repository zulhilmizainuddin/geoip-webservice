package models

type Geolocation struct {
	IPAddress string	`json:"ipaddress"`
	City string      	`json:"city"`
	Country string   	`json:"country"`
	Latitude float32 	`json:"latitude"`
	Longitude float32	`json:"longitude"`
}