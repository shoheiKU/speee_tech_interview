FROM golang:1.21rc2-alpine3.18
RUN apk update && apk add git

WORKDIR /go/app