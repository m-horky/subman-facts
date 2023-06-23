.PHONY: help
help:
	@echo "build .......... compile the package"
	@echo "run ............ compile and run the package"

.PHONY: build
build:
	VERSION=`git rev-parse --short HEAD`; \
	go build -ldflags="-X main.version=$$VERSION" -o ./bin/subman-facts ./cmd/subman-facts

.PHONY: run
run: build
	./bin/subman-facts
