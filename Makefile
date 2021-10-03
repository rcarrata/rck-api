all: build tag push
test: lint unit-test

NAME=rck
VERSION=0.2
REGISTRY="quay.io/rcarrata"
TOOL="docker"

install:
	@go build .

build: 
	@go version
	@${TOOL} build -t localhost/in-cluster:${VERSION} .
	
tag:
	@${TOOL} tag localhost/in-cluster:${VERSION} ${REGISTRY}/${NAME}:${VERSION}

push: 
	@${TOOL} push ${REGISTRY}/${NAME}:${VERSION}
