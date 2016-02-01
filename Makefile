ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${ROOT}:${GOPATH}
export GOPATH := ${GOPATH}


test:  test-style test-trans test-svg

test-svg:
	@go test ./

test-style:
	@cd ./style & go test ./style

test-trans:
	@cd ./transform & go test ./transform

