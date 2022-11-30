FROM golang:1.19.3-alpine as dev

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

RUN apk upgrade --update && apk --no-cache add git
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

ARG CGO_ENABLED=0
ARG GOOS=linux
# アーキテクチャを指定する
# ARG GOARCH=amd64
ARG GOARCH=arm64
