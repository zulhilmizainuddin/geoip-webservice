FROM golang:alpine

RUN apk add --update git wget unzip nodejs && \
    npm install -g pm2

WORKDIR /go/src/geoip-webservice
COPY . .

RUN go get github.com/gin-gonic/gin && \
    go get github.com/ip2location/ip2location-go && \
    go get github.com/spf13/viper

ENV DATABASE_ZIP IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP

RUN wget https://www.dropbox.com/s/3dimi97x9f0qwxe/IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP?dl=1 -O ./databases/ip2location/$DATABASE_ZIP && \
    unzip ./databases/ip2location/$DATABASE_ZIP -d ./databases/ip2location && \
    rm -fv ./databases/ip2location/$DATABASE_ZIP

EXPOSE 4000

CMD ["pm2-docker", "start", "go", "--", "run", "/go/src/geoip-webservice/main.go"]
