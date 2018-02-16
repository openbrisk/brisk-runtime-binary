# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-binary
COPY  ./src/*.go /go/src/github.com/openbrisk/brisk-runtime-binary

RUN GOOS=linux GOARCH=amd64 go build -o server .

# Release stage
# NOTE: Keep this up to date.
FROM ubuntu:18.04 

WORKDIR /app

RUN apt-get update && apt-get install -y jq

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-binary/server /app/server
COPY startup.sh .
COPY function-wrapper.sh .

RUN chmod +x startup.sh && chmod +x function-wrapper.sh

EXPOSE 8080
ENTRYPOINT [ "./startup.sh" ]