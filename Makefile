.PHONY: build test lint install clean

build:
	go build -o bin/addigy-cli ./cmd/addigy-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/addigy-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/addigy-mcp ./cmd/addigy-mcp

install-mcp:
	go install ./cmd/addigy-mcp

build-all: build build-mcp
