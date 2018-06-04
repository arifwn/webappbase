SHELL := /bin/bash

PKG := github.com/arifwn/webappbase

TARGET := ./_build/webserver

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

SRC_DIRS := cmd internal pkg # directories which hold app source (not vendored)

.PHONY: all build clean install uninstall fmt simplify check run

all: build

build:
	@go build -o ${TARGET} "${PKG}/cmd/webserver"

clean:
	@true

install:
	@true

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

run:
	@go run "${PKG}/cmd/webserver/main.go"

watch:
	# https://github.com/canthefason/go-watcher
	cd cmd/webserver/; watcher -watch ${PKG}

test:
	@hack/test.sh $(SRC_DIRS)
