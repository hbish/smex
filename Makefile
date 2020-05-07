NAME=smex
VERSION := $(shell cat ./VERSION)
PKGS := $(shell go list ./... | grep -v vendor)

BIN_DIR := $(CURDIR)/bin
BUILD_DIR := $(CURDIR)/build

LINT_BIN := $(GOPATH)/bin/golangci-lint

prepare-lint: # prepare lint dependency
	@if [ -z `which $(BIN_DIR)/golangci-lint` ]; then \
		echo "[downloading] installing golangci-lint";\
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) v1.26.0;\
	fi

prepare-cov:
	@if [ -z `which $(BIN_DIR)/gocov` ]; then \
		echo "[go get] installing gocov";\
		GO111MODULE=off GOBIN=$(BIN_DIR) go get github.com/axw/gocov/gocov;\
	fi

prepare: prepare-lint prepare-cov # prepare ci dependency
	@if [ -z `which $(BIN_DIR)/gox` ]; then \
		echo "[go get] installing gox";\
		GO111MODULE=off GOBIN=$(BIN_DIR) go get github.com/mitchellh/gox;\
	fi
	@if [ -z `which $(BIN_DIR)/ghr` ]; then \
		echo "[go get] installing ghr";\
		GO111MODULE=off GOBIN=$(BIN_DIR) go get github.com/tcnksm/ghr;\
	fi
	@if [ -z `which $(BIN_DIR)/semantics` ]; then \
		echo "[go get] installing semantics";\
		GO111MODULE=off GOBIN=$(BIN_DIR) go get github.com/stevenmatthewt/semantics;\
	fi

init: mod # init repo for local development
	git config core.hooksPath .githooks
.PHONY: init

build-ci: prepare build dist # run ci build

build: mod # build and compile smex excutables
	@rm -rf build/
	@$(BIN_DIR)/gox -ldflags "-X main.Version=$(VERSION)" \
	-osarch="darwin/amd64" \
	-osarch="linux/386" \
	-osarch="linux/amd64" \
	-osarch="windows/amd64" \
	-osarch="windows/386" \
	-output "build/{{.Dir}}_$(VERSION)_{{.OS}}_{{.Arch}}/$(NAME)" \
	${PKGS}
.PHONY: build

dist: # prepare for distribution
	$(eval FILES := $(shell ls build))
	@rm -rf dist && mkdir dist
	@for f in $(FILES); do \
		(cd $(shell pwd)/build/$$f && tar -cvzf ../../dist/$$f.tar.gz *); \
		(cd $(shell pwd)/dist && shasum -a 512 $$f.tar.gz > $$f.sha512); \
		echo $$f; \
	done

clean: # remove build artifacts
	@rm -rf $(BUILD_DIR)
.PHONY: clean

mod: clean # download go modules
	@go mod download
	@go mod tidy

lint: prepare-lint # check for errors in code
	@$(BIN_DIR)/golangci-lint run -p format -p unused -p bugs
.PHONY: lint

test: prepare-cov # run unit tests
	@$(BIN_DIR)/gocov test $(PKGS) | $(BIN_DIR)/gocov report
.PHONY: test

release-ci: dist # run ci release
	$(eval TAG := $(shell $(BIN_DIR)/semantics -output-tag))
	if [ "$(TAG)" ]; then \
	  $(BIN_DIR)/ghr -t $(GITHUB_TOKEN) -u $(CIRCLE_PROJECT_USERNAME) -r $(CIRCLE_PROJECT_REPONAME) --replace $(TAG) dist/; \
	fi
.PHONY: release-ci

help: Makefiles
	@echo
	@echo " Choose a command run in "$(NAME)":"
	@echo
	@cat Makefile | grep '^[^ ]*:' | grep -v '.PHONY' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t
.PHONY: help
