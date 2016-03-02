ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${ROOT}:${GOPATH}
export GOPATH := ${GOPATH}


build:
	(cd ./tools && make && cd ../ && ./tools/gen `pwd` && cd ./tools && make clean)


test:  test-style test-trans test-svg

test-svg:
	@go test ./

test-style:
	@cd ./style & go test ./style

test-trans:
	@cd ./transform & go test ./transform


# reinstall
reinstall:
	@git push
	@go get -u github.com/iostrovok/svg/style
	@go get -u github.com/iostrovok/svg/transform
	@go get -u github.com/iostrovok/svg
