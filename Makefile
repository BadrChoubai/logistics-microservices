include .env

BINS ?= gateway shipment telemetry
BUILD_DIRS := bin

# Used internally.  Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell GOTOOLCHAIN=local go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell GOTOOLCHAIN=local go env GOARCH))
GO := $(if $(GOVERSION),$(GOVERSION),$(shell GOTOOLCHAIN=local go env GOVERSION))
DOCKER := podman

SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
GOFLAGS ?=
VERSION ?= $(shell git describe --tags --always --dirty)


all: # @HELP build container image 
all: build

build: # @HELP build app for local development
build: deps ci $(BUILD_DIRS) $(addprefix build-,$(BINS))


$(BUILD_DIRS):
	mkdir -p $@


$(addprefix build-,$(BINS)): build-%:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$* $(GOFLAGS) ./cmd/$*/main.go


ci: # @HELP ci steps(lint, test)
ci: lint test


clean: # @HELP clean build artifacts
clean: image-clean
	rm -rf ./bin


docs: # @HELP generate Swagger API documentation
docs: $(addprefix docs-,$(BINS))
	swag fmt


$(addprefix docs-,$(BINS)): docs-%:
	swag init --generalInfo main.go \
		--dir ./cmd/$*,./internal/$*/handler \
		--parseInternal \
		-o api/swagger/$*


deps: # @HELP go mod tidy, download
deps:
	go mod tidy
	go mod download


image: # @HELP build all container images
image: $(addprefix image-,$(BINS))


$(addprefix image-,$(BINS)): image-%:
	@if $(DOCKER) image inspect logistics-$*:$(VERSION) >/dev/null 2>&1; then \
		echo "Image logistics-$*:$(VERSION) already exists"; \
	else \
		$(DOCKER) build --build-arg SERVICE=$* -t logistics-$*:$(VERSION) -f ./manifests/$*/Dockerfile .; \
	fi


image-clean: # @HELP remove all built images
image-clean:
	$(foreach bin, $(BINS), $(DOCKER) rmi logistics-$(bin):$(VERSION) || true;)


lint: # @HELP lint with golangci-lint
lint:
	golangci-lint run

test: # @HELP runs unit tests
test:
	go test ./...

version: # @HELP prints the version string
version:
	@echo $(VERSION)


help: # @HELP prints this message
help:
	echo "VARIABLES:"
	echo "  BINS = $(BINS)"
	echo "  OS = $(OS)"
	echo "  ARCH = $(ARCH)"
	echo "  GOFLAGS = $(GOFLAGS)"
	echo "  GO = $(GO)"
	echo "  DOCKER = $(DOCKER)"
	echo
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '

.SILENT: help
.PHONY: all build ci clean deps docs image image-clean lint test version help \
	$(addprefix build-,$(BINS)) \
	$(addprefix image-,$(BINS)) \
	$(addprefix docs-,$(BINS))
