.PHONY: build

build:
	docker build -t brisk-runtime-binary .

run:
	docker run -it \
	-p 8080:8080 \
	-e MODULE_NAME=figlet \
	-e FUNCTION_DEPENDENCIES=figlet.deps \
	-e FUNCTION_HANDLER=figlet \
	-e FUNCTION_TIMEOUT=25 \
	brisk-runtime-binary

compile:
	go build ./src/server.go