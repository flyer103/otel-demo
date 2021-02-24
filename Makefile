GODIR = $(shell go list ./... | grep -v /vendor/)
PKG := github.com/flyer103/otel-demo
BUILD_IMAGE ?= golang:1.16
GOARCH := amd64
# GOOS := linux
GOOS := darwin
VERSION := $(shell git rev-parse HEAD)

LDFLAGS_OTEL_DEMO := -ldflags "-X ${PKG}/pkg/version.Version=${VERSION}"

pre-build:
	@echo "pre build"
	@echo "clean all flycheck files"
	@find . -name "flycheck*" | xargs rm -f
.PHONY: pre-build

builds: build-demo
.PHONY: builds

build-dirs: pre-build
	@mkdir -p .go/src/$(PKG) .go/bin .cache
	@mkdir -p release
.PHONY: build-dirs

build-demo: build-dirs
	@docker run                                                            \
	    --rm                                                               \
	    -ti                                                                \
	    -u $$(id -u):$$(id -g)                                             \
	    -v $$(pwd)/.go:/go                                                 \
	    -v $$(pwd):/go/src/$(PKG)                                          \
	    -v $$(pwd)/release:/go/bin                                         \
	    -v $$(pwd)/.cache:/.cache            			       \
	    -e GOOS=$(GOOS)                                                    \
	    -e GOARCH=$(GOARCH)                                                \
	    -e CGO_ENABLED=0                                                   \
            -e GO111MODULE=off                                                 \
	    -w /go/src/$(PKG)                                                  \
	    $(BUILD_IMAGE)                                                     \
	    go build -o ./release/demo $(LDFLAGS_OTEL_DEMO) ./cmd/
.PHONY: build-demo

