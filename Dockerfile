FROM golang:alpine

RUN apk add --update git nodejs && \
    npm install -g pm2

WORKDIR /go/src/geoip-webservice
COPY . /go/src/geoip-webservice

RUN go get github.com/gin-gonic/gin && \
    go get github.com/ip2location/ip2location-go && \
    go get github.com/spf13/viper

EXPOSE 4000

CMD ["pm2-docker", "start", "go", "--", "run", "/go/src/geoip-webservice/main.go"]