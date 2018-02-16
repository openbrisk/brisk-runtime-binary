# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-binary
COPY  ./src/*.go /go/src/github.com/openbrisk/brisk-runtime-binary

RUN GOOS=linux GOARCH=amd64 go build -o server .

# Release stage
# NOTE: Keep this up to date.
FROM ubuntu:17:10 

WORKDIR /app

RUN apt-get update

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-binary/server /app/server
COPY startup.sh .
COPY function-wrapper.sh .

EXPOSE 8080
ENTRYPOINT [ "./startup.sh" ]