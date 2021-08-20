all: bin/example
test: lint unit-test

NAME=in-cluster
VERSION=0.6
REGISTRY="quay.io/rcarrata"

build: 
	@go version
	@podman build -t localhost/in-cluster:${VERSION} .
	
tag:
	@podman tag localhost/in-cluster:${VERSION} ${REGISTRY}/${NAME}:${VERSION}

push: 
	@podman push ${REGISTRY}/${NAME}:${VERSION}
