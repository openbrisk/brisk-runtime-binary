# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-binary
COPY  ./src/*.go /go/src/github.com/openbrisk/brisk-runtime-binary

RUN GOOS=linux GOARCH=amd64 go build -o server .

# Release stage
FROM ubuntu:latest

WORKDIR /app

RUN apt-get update

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-binary/server /app/server

EXPOSE 8080
ENTRYPOINT [ "./server" ]