.PHONY: build run compile start

build:
	docker build -t brisk-runtime-binary .

run:
	docker run -it \
	-p 8080:8080 \
	-e MODULE_NAME=figlet \
	-e FUNCTION_HANDLER=execute \
	-e FUNCTION_TIMEOUT=10 \
	-v `pwd`/examples:/openbrisk \
	brisk-runtime-binary

compile:
	go build ./src/server.go

start:
	export MODULE_NAME=figlet && \
	export FUNCTION_HANDLER=execute && \
	export FUNCTION_TIMEOUT=10 && \
	./server