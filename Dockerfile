# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-binary
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server ./src/server.go

# Release stage
FROM ubuntu:18.04 

WORKDIR /app

RUN apt-get update && apt-get install -y jq jo

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-binary/server /app/server
COPY startup.sh .
COPY function-wrapper.sh .
COPY forward.sh .

RUN chmod +x startup.sh && chmod +x function-wrapper.sh && chmod +x forward.sh

EXPOSE 8080
ENTRYPOINT [ "./startup.sh" ]