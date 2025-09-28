BINARY=bin/server

.PHONY: help run-dev run-prod build-dev build-prod clean

help:
	@echo "Available targets: run-dev run-prod build-dev build-prod clean"

run-dev:
	@echo "Running (dev)..."
	APP_ENV=dev go run ./cmd/server

run-prod:
	@echo "Running (prod)..."
	APP_ENV=prod go run ./cmd/server

build-dev:
	@echo "Building dev binary..."
	mkdir -p bin
	APP_ENV=dev go build -o $(BINARY)-dev ./cmd/server

build-prod:
	@echo "Building prod binary..."
	mkdir -p bin
	APP_ENV=prod go build -o $(BINARY)-prod ./cmd/server

clean:
	@echo "Cleaning..."
	rm -rf bin
