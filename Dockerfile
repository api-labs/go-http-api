FROM golang:1.14

WORKDIR /usr/src/go-http-api
COPY . .

RUN go build -o go-http-api