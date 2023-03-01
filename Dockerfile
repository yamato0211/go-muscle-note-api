FROM golang:1.19.3-alpine as dev

RUN mkdir -p /go/src/api
COPY . /go/src/api
WORKDIR /go/src/api

RUN go install

RUN go build -o /go/bin/air github.com/cosmtrek/air
