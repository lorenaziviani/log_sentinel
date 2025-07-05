# Makefile for log_sentinel/collector

.PHONY: all test lint run integration-test

all: test lint

# Unit tests
unit-test:
	go test -v ./...

test: unit-test

# Lint (using go vet)
lint:
	go vet ./...

# Run the collector
run:
	go run main.go

# Integration test (sends real logs)
integration-test:
	bash ../logs-test.sh 