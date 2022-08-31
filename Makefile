.POSIX:
SHELL := /bin/bash
GO := $(shell which go)
BIN_NAME := akita
# Build the application
.PHONY: build
build: fmt
	@echo -e "\033[32mBuilding the application...\033[0m"
	$(GO) build -ldflags "-s -w" -o ./dist/$(BIN_NAME) ./cmd/$(BIN_NAME)/main.go
	@echo -e "\033[32mBuild finished.\033[0m"
# Local development
.PHONY: dev
dev:
	$(GO) build -o ./dist/$(BIN_NAME)-dev ./cmd/$(BIN_NAME)/main.go
	./dist/$(BIN_NAME)-dev serve
# Format the code
.PHONY: fmt
fmt:
	@echo -e "\033[32mFormatting the code...\033[0m"
	$(GO) fmt ./...
	@echo -e "\033[32mFormatting finished.\033[0m"
# Install dependencies
.PHONY: install
install:
	@echo -e "\033[32mInstalling dependencies...\033[0m"
	$(GO) mod download
	@echo -e "\033[32mDependencies installed.\033[0m"
# Lint the code
.PHONY: lint
lint:
	@echo -e "\033[32mLinting the code...\033[0m"
	$(GO) vet ./...
	@echo -e "\033[32mLinting finished.\033[0m"
