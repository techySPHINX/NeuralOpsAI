
.PHONY: help dev-tools lint test build docker proto kind-up temporal-up run-gateway

help:
	@echo "Makefile for NeuralOps"
	@echo ""
	@echo "Targets:"
	@echo "  dev-tools      - Install development tools"
	@echo "  lint           - Run linters"
	@echo "  test           - Run tests"
	@echo "  build          - Build binaries"
	@echo "  docker         - Build docker images"
	@echo "  proto          - Generate protobuf files"
	@echo "  kind-up        - Start a local kind cluster"
	@echo "  temporal-up    - Start a local temporal server"
	@echo "  run-gateway    - Run the API gateway"

dev-tools:
	@echo "Installing development tools..."
	# Add commands to install tools like buf, helm, etc.

lint:
	@echo "Linting..."
	# Add linting commands

test:
	@echo "Testing..."
	# Add test commands

build:
	@echo "Building..."
	# Add build commands

docker:
	@echo "Building docker images..."
	# Add docker build commands

proto:
	@echo "Generating protobuf files..."
	# Add protobuf generation commands

kind-up:
	@echo "Starting kind cluster..."
	# Add kind cluster startup commands

temporal-up:
	@echo "Starting temporal server..."
	# Add temporal startup commands

run-gateway:
	@echo "Running API gateway..."
	# Add command to run the api-gateway
