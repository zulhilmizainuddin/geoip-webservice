# geoip-webservice ![Build Status](https://github.com/zulhilmizainuddin/geoip-webservice/actions/workflows/build.yml/badge.svg?branch=master)

RESTful service for retrieving geolocations from IP addresses using Go and Gin framework

Geolocation data provided by IP2Location LITE and MaxMind GeoLite2 databases.

## Getting Started

Download IP2LOCATION-LITE-DB5.IPV6.BIN from https://lite.ip2location.com/database-ip-country-region-city-latitude-longitude and place into the databases/ip2location directory.

Download GeoLite2-City_xxxxxxxx.tar.gz from http://dev.maxmind.com/geoip/geoip2/geolite2 and place into the databases/maxmind directory.

Start the service.
```bash
go run main.go
```

## Endpoint
### IP2Location
```
http://localhost:4000/ip2location?ipaddress={ip-address}
```

Make a HTTP GET request to the endpoint to retrieve the IP address geolocation. Replace {ip-address} with the IP address to be searched.

## Usage Example
```bash
curl --location 'http://localhost:4000/ip2location?ipaddress=123.123.123.123'
```

### Result Example
```javascript
{
    "ipaddress":"123.123.123.123",
    "city":"Beijing",
    "country":"China",
    "latitude":39.9289,
    "longitude":116.3883
}
```

### Error Example
```javascript
{
    "error_message":"Invalid IP address provided"
}
```
