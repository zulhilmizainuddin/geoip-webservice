FROM golang:alpine

RUN apk add --update git wget unzip tar nodejs && \
    npm install -g pm2

WORKDIR /go/src/geoip-webservice
COPY . .

RUN go get github.com/gin-gonic/gin && \
    go get github.com/ip2location/ip2location-go && \
    go get github.com/spf13/viper && \
    go get github.com/oschwald/geoip2-golang

ENV IP2LOCATION_DATABASE IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP
ENV MAXMIND_DATABASE GeoLite2-City.tar.gz

RUN wget https://www.dropbox.com/s/3dimi97x9f0qwxe/IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP?dl=1 -O ./databases/ip2location/$IP2LOCATION_DATABASE && \
    unzip -o ./databases/ip2location/$IP2LOCATION_DATABASE -d ./databases/ip2location && \
    rm -fv ./databases/ip2location/$IP2LOCATION_DATABASE

RUN wget https://www.dropbox.com/s/4zdyyip84p5opl7/GeoLite2-City.tar.gz?dl=1 -O ./databases/maxmind/$MAXMIND_DATABASE && \
    tar -xvzf ./databases/maxmind/$MAXMIND_DATABASE --strip-components=1 -C ./databases/maxmind && \
    rm -fv ./databases/maxmind/$MAXMIND_DATABASE

EXPOSE 4000

CMD ["pm2-docker", "start", "go", "--", "run", "/go/src/geoip-webservice/main.go"]
