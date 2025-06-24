.PHONY: test test-coverage test-ci lint fmt clean build wire

# Test commands
test:
	go test -v -race ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-ci:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Lint and format
lint:
	golangci-lint run

fmt:
	go fmt ./...
	gofmt -s -w .

# Build commands
build:
	go build -v -o hakushi .

wire:
	wire ./internal/wire

# Clean
clean:
	rm -f hakushi coverage.out coverage.html
	go clean -testcache

# Docker commands for testing
docker-test:
	docker-compose -f ../docker/test/compose.yaml up --build --abort-on-container-exit
	docker-compose -f ../docker/test/compose.yaml down

docker-test-down:
	docker-compose -f ../docker/test/compose.yaml down -v