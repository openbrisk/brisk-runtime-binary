.PHONY: build run compile start

build:
	docker build -t brisk-runtime-binary .

run:
	docker run -it \
	-p 8080:8080 \
	-e MODULE_NAME=figlet \
	-e FUNCTION_HANDLER=figlet \
	-e FUNCTION_TIMEOUT=10 \
	brisk-runtime-binary

compile:
	go build ./src/server.go

start:
	MODULE_NAME=hello_world \
	FUNCTION_HANDLER=execute \
	./function-wrapper.sh