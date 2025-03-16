.PHONY: all test lint format 

# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

all: lint test build

test:
	$(GOTEST) -v -race -cover ./...

lint:
	golangci-lint run --config .golangci.yml

format:
	golangci-lint run --fix --timeout=5m --config .golangci.yml
	$(GOMOD) tidy
