FROM golang:1.14.0-alpine3.11

WORKDIR /go/app
RUN apk update && apk upgrade && apk add \
  git \
  tzdata
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && apk del tzdata && rm -rf /var/cache/apk/*

RUN go get github.com/pilu/fresh
RUN go get github.com/gorilla/websocket
