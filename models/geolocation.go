package models

type Geolocation struct {
	IPAddress string	`json:"ipaddress"`
	City string      	`json:"city"`
	Country string   	`json:"country"`
	Latitude float32 	`json:"latitude"`
	Longitude float32	`json:"longitude"`
}

type Error struct {
	ErrorMessage string	`json:"error_message"`
}

type Retriever interface {
	Query(ipaddress string) interface{}
}