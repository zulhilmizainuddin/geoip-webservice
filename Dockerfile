FROM golang:1.20.3-alpine

RUN apk add --update git wget unzip tar

WORKDIR /go/src/geoip-webservice
COPY . .

ENV IP2LOCATION_DATABASE IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP
ENV MAXMIND_DATABASE GeoLite2-City.tar.gz

RUN wget https://www.dropbox.com/s/3dimi97x9f0qwxe/IP2LOCATION-LITE-DB5.IPV6.BIN.ZIP?dl=1 -O ./databases/ip2location/$IP2LOCATION_DATABASE && \
    unzip -o ./databases/ip2location/$IP2LOCATION_DATABASE -d ./databases/ip2location && \
    rm -fv ./databases/ip2location/$IP2LOCATION_DATABASE

RUN wget https://www.dropbox.com/s/4zdyyip84p5opl7/GeoLite2-City.tar.gz?dl=1 -O ./databases/maxmind/$MAXMIND_DATABASE && \
    tar -xvzf ./databases/maxmind/$MAXMIND_DATABASE --strip-components=1 -C ./databases/maxmind && \
    rm -fv ./databases/maxmind/$MAXMIND_DATABASE

RUN go build -o /go/bin/geoip-webservice

EXPOSE 4000

CMD ["/go/bin/geoip-webservice"]