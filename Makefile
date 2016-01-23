ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${ROOT}:${GOPATH}
export GOPATH := ${GOPATH}


test: test-svg test-style

test-svg:
	@go test ./

test-style:
	@cd ./style & go test ./style

