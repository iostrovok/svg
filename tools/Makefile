ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${ROOT}:${GOPATH}
export GOPATH := ${GOPATH}

.PHONY: all


all: path clean fmt build 

path:
	@echo ${GOPATH}

build:
	@go build -o gen ./gen.go
	@echo "Done"

fmt:
	@go fmt ./...

clean:
	@go clean -i ./...

