FROM golang:1.13.4-alpine3.10 AS builder

WORKDIR /go/src/github.com/mollylucky/dog-park

COPY . .

RUN apk add git \
        && go get -u -v -d ./... \
        && go build -o dog-park ./

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /go/src/github.com/mollylucky/dog-park/dog-park dog-park

ENTRYPOINT [ "dog-park" ]