.PHONY: help
help:
	@echo "build .......... compile the package"

.PHONY: build
build:
	VERSION=`git rev-parse --short HEAD`; \
	go build -ldflags="-X main.version=$$VERSION" -o bin/subman-facts ./cmd/subman-facts
