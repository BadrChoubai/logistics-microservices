include .env

BINS ?= gateway

# Used internally.  Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell GOTOOLCHAIN=local go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell GOTOOLCHAIN=local go env GOARCH))
GO := $(if $(GOVERSION),$(GOVERSION),$(shell GOTOOLCHAIN=local go env GOVERSION))

SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
GOFLAGS ?=
VERSION ?= $(shell git describe --tags --always --dirty)


all: # @HELP build container image 
all: build

build: # @HELP build app for local development
build: deps ci
	mkdir -p bin
	$(foreach bin, $(BINS), \
		GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$(bin) $(GOFLAGS) ./cmd/$(bin)/main.go;)

ci: # @HELP ci steps(lint, test)
ci: lint test

$(addprefix build-,$(BINS)): build-%:
	mkdir -p bin
	GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$* $(GOFLAGS) ./cmd/$*/main.go

clean: # @HELP clean build artifacts
clean: image-clean
	rm -rf ./bin

docs: # @HELP generate documentation
docs:
	swag fmt
	swag init --generalInfo cmd/gateway/main.go --dir . --parseInternal -o api/swagger

deps: # @HELP go mod tidy, download
deps:
	go mod tidy
	go mod download

image: # @HELP build all docker images
image:
	$(foreach bin, $(BINS), \
		docker image inspect logistics-$(bin):$(VERSION) >/dev/null 2>&1 	\
		&& echo "Image logistics-$(bin):$(VERSION) already exists"			\
		|| docker build --build-arg SERVICE=$(bin) -t logistics-$(bin):$(VERSION) -f ./manifests/Dockerfile .;)

$(addprefix image-,$(BINS)): image-%:
	docker image inspect logistics-$*:$(VERSION) >/dev/null 2>&1 	\
	&& echo "Image logistics-$*:$(VERSION) already exists"			\
	|| docker build --build-arg SERVICE=$* -t logistics-$*:$(VERSION) -f ./manifests/Dockerfile .


image-clean: # @HELP remove all built images
image-clean:
	$(foreach bin, $(BINS), docker rmi logistics-$(bin):$(VERSION) || true;)


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
	$(addprefix image-,$(BINS))
