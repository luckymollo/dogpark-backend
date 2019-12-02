FROM golang:1.13.4-alpine3.10 AS builder

WORKDIR /go/src/github.com/mollylucky/dog-park

COPY . .

RUN apk add git \
        && go get "github.com/stretchr/testify/assert" \
        && go get "github.com/gin-gonic/gin" \
        && go build -o dog-park ./

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /go/src/github.com/mollylucky/dog-park/dog-park dog-park

ENTRYPOINT [ "dog-park" ]