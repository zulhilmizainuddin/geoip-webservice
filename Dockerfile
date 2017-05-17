FROM golang:1.8

RUN apt-get update
RUN apt-get install -y nodejs
RUN apt-get install -y npm

RUN npm install -g pm2

RUN ln -s $(which nodejs) /usr/bin/node

WORKDIR /go/src/geoip-webservice
COPY . /go/src/geoip-webservice

RUN go get github.com/gin-gonic/gin
RUN go get github.com/ip2location/ip2location-go
RUN go get github.com/spf13/viper

EXPOSE 4000

CMD ["pm2-docker", "start", "go", "--", "run", "/go/src/geoip-webservice/main.go"]