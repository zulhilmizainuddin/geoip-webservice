# geoip-webservice [![Build Status](https://travis-ci.org/zulhilmizainuddin/geoip-webservice.svg?branch=master)](https://travis-ci.org/zulhilmizainuddin/geoip-webservice)
RESTful service for retrieving geolocations from IP addresses using Go and Gin framework

Geolocation data provided by IP2Location LITE and MaxMind GeoLite2 databases.

## Getting Started

Install Go. Follow the instruction at https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04.

Get project dependencies.
```bash
go get -v github.com/gin-gonic/gin
go get -v github.com/ip2location/ip2location-go
go get -v github.com/spf13/viper
go get -v github.com/oschwald/geoip2-golang
```

Install Node.js. Follow the instruction at https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-16-04.

Install PM2.
```bash
npm install -g pm2
```

Download IP2LOCATION-LITE-DB5.IPV6.BIN from https://lite.ip2location.com/database-ip-country-region-city-latitude-longitude and place into the databases/ip2location directory.

Download GeoLite2-City_xxxxxxxx.tar.gz from http://dev.maxmind.com/geoip/geoip2/geolite2 and place into the databases/maxmind directory.

Start the service using PM2.
```bash
pm2 start --name geoip-webservice go -- run main.go
```


## Endpoint
### IP2Location
```
http://localhost:4000/ip2location?ipaddress={ip-address}
```

Make a HTTP GET request to the endpoint to retrieve the IP address geolocation. Replace {ip-address} with the IP address to be searched.

## Usage Example
```bash
curl http://localhost:4000/ip2location?ipaddress=123.123.123.123
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
