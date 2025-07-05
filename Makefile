# Makefile for log_sentinel/collector

.PHONY: all test lint run integration-test train-ml predict-ml test-ml

all: test lint

# Unit tests
unit-test:
	go test -v ./...

test: unit-test

# Lint (using go vet)
lint:
	go vet ./...

# Lint (using golangci-lint)
lint:
	golangci-lint run ./...

# Run the collector
run:
	go run main.go

# Integration test (sends real logs)
integration-test:
	bash ./cmd/collector/logs-test.sh

# Train the ML model
train-ml:
	bash ./cmd/ml/train-ml.sh

# Test the ML /predict endpoint
predict-ml:
	bash ./cmd/ml/predict-ml.sh

# Unit tests ML (Python)
test-ml:
	cd cmd/ml && pytest 