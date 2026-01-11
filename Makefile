BINARY=bin/server
VERSION := $(shell cat .VERSION)

.PHONY: help run-dev run-prod build-dev build-prod build-prod-docker clean

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

build-dev-docker:
	@echo "Building dev image" $(VERSION)
	docker build -t "google-takeout-sucks/google_takeout_sucks_auth:dev-$(VERSION)" -f Dockerfile.dev .

build-prod-docker:
	@echo "Building prod image $(VERSION)"
	docker build -t "google-takeout-sucks/google_takeout_sucks_auth:$(VERSION)" -f Dockerfile.prod .

push-prod-docker:
	@echo "Pushing prod docker image"
	docker tag google-takeout-sucks/google_takeout_sucks_auth:$(VERSION) us-west1-docker.pkg.dev/download-photos-417323/google-takeout-sucks/google_takeout_sucks_auth:$(VERSION)
	docker push us-west1-docker.pkg.dev/download-photos-417323/google-takeout-sucks/google_takeout_sucks_auth:$(VERSION)

clean:
	@echo "Cleaning..."
	rm -rf bin
