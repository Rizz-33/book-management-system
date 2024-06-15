# Makefile for a Go project

# Define variables
APP_NAME := book-management-system
SRC_DIR := ./src
BIN_DIR := ./bin

# Default target
all: build

# Build the application
build:
    @echo "Building the application..."
    go build -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)/main.go

# Clean the build artifacts
clean:
    @echo "Cleaning build artifacts..."
    rm -rf $(BIN_DIR)/*

# Run the application
run: build
    @echo "Running the application..."
    $(BIN_DIR)/$(APP_NAME)

# Run tests
test:
    @echo "Running tests..."
    go test ./...

# Help message
help:
    @echo "Usage: make [target]"
    @echo "Targets:"
    @echo "  all      - Build the application"
    @echo "  build    - Build the application"
    @echo "  clean    - Clean the build artifacts"
    @echo "  run      - Run the application"
    @echo "  test     - Run tests"
    @echo "  help     - Display this help message"
