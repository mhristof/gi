#
# vim:ft=make
#

MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
.DELETE_ON_ERROR:
.ONESHELL:

GIT_TAG := $(shell git describe --tags --always --dirty=+)


.PHONY: all
all: ./bin/gi.darwin

./bin/gi.%: $(shell find ./ -name '*.go')
	GOOS=$* go build -o $@ -ldflags "-X github.com/mhristof/gi/cmd.version=$(GIT_TAG)" main.go

.PHONY: fast-test
fast-test:  ## Run fast tests
	go test ./... -tags fast

.PHONY: test
test:	## Run all tests
	go test ./...

.PHONY: clean
clean:
	rm -rf bin/gi.*

.PHONY: help
help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.*## /:/g' | column -t -s:

install: ./bin/gi.darwin
	rm ~/.local/bin/gi # macos permissions are bound to the current executable
	cp ./bin/gi.darwin ~/.local/bin/gi
