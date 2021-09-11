all: install build tag push
test: lint unit-test

NAME=in-cluster
VERSION=0.8
REGISTRY="quay.io/rcarrata"

install:
	@go build .

build: 
	@go version
	@podman build -t localhost/in-cluster:${VERSION} .
	
tag:
	@podman tag localhost/in-cluster:${VERSION} ${REGISTRY}/${NAME}:${VERSION}

push: 
	@podman push ${REGISTRY}/${NAME}:${VERSION}
