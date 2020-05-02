NAME=smex
VERSION := $(shell git describe --tags --always --dirty)
PKGS := $(shell go list ./... | grep -v vendor)

BIN_DIR := $(CURDIR)/bin

LINT_BIN := $(GOPATH)/bin/golangci-lint

$(LINT_BIN):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint


init:
	git config core.hooksPath .githooks
.PHONY: init

## build: compile and builds smex excutables
build:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME)-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME)-linux-amd64 main.go
.PHONY: build

## clean: remove build artifacts
clean:
	rm -rf bin
.PHONY: clean

lint: $(LINT_BIN)
	@$(LINT_BIN) run -p format -p unused -p bugs
.PHONY: lint

test:
	go test -v -cover $(PKGS)
.PHONY: test

help: Makefile
	@echo
	@echo " Choose a command run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
.PHONY: help
