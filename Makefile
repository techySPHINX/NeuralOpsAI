
.PHONY: help dev-tools lint test build docker proto kind-up temporal-up run-gateway dashboard build-services

BUF_VERSION=1.28.1
GOBIN=$(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN=$(shell go env GOPATH)/bin
endif

help:
	@echo "Makefile for NeuralOps"
	@echo ""
	@echo "Targets:"
	@echo "  dev-tools      - Install development tools"
	@echo "  lint           - Run linters"
	@echo "  test           - Run tests"
	@echo "  build          - Build all binaries"
	@echo "  build-services - Build all Go services"
	@echo "  dashboard      - Build and run dashboard"
	@echo "  docker         - Build docker images"
	@echo "  proto          - Generate protobuf files"
	@echo "  kind-up        - Start a local kind cluster"
	@echo "  temporal-up    - Start a local temporal server"
	@echo "  run-gateway    - Run the API gateway"
	@echo "  up             - Start all services with docker-compose"
	@echo "  down           - Stop all services"

dev-tools:
	@echo "Installing development tools..."
	@go install github.com/bufbuild/buf/cmd/buf@v$(BUF_VERSION)
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

lint:
	@echo "Linting..."
	@golangci-lint run ./...

test:
	@echo "Testing..."
	@go test -v ./...

build: build-services dashboard
	@echo "Build complete!"

build-services:
	@echo "Building Go services..."
	@go build -o bin/api-gateway ./cmd/api-gateway
	@go build -o bin/orchestrator ./cmd/orchestrator
	@go build -o bin/ai-engine ./cmd/ai-engine
	@go build -o bin/iceberg-manager ./cmd/iceberg-manager
	@go build -o bin/model-registry ./cmd/model-registry
	@go build -o bin/automl-service ./cmd/automl-service
	@echo "Services built successfully!"

dashboard:
	@echo "Building dashboard..."
	@cd web/dashboard && npm install && npm run build
	@echo "Dashboard built successfully!"

docker:
	@echo "Building docker images..."
	@docker-compose -f docker-compose.enhanced.yaml build

proto:
	@echo "Generating protobuf files..."
	@cd api/proto && buf generate

kind-up:
	@echo "Starting kind cluster..."
	@kind create cluster --name neuralops

temporal-up:
	@echo "Starting temporal server..."
	@docker run -d -p 7233:7233 temporalio/auto-setup:latest

run-gateway:
	@echo "Running API gateway..."
	@go run ./cmd/api-gateway

up:
	@echo "Starting all services..."
	@docker-compose -f docker-compose.enhanced.yaml up -d
	@echo "Services started! Dashboard: http://localhost:8080"

down:
	@echo "Stopping all services..."
	@docker-compose -f docker-compose.enhanced.yaml down

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf web/dashboard/dist/
	@echo "Clean complete!"
