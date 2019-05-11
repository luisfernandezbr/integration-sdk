#
# Makefile for building all things related to this repo
#
NAME := integration-sdk
ORG := pinpt
PKG := $(ORG)/$(NAME)
SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})

VERSION ?= $(shell git describe --always --dirty='-dev')
BUILD ?= $(shell git rev-parse HEAD | cut -c1-8)
COMMITSHA ?= $(shell git rev-parse HEAD)

all: help

generate:
	@cd $(GOPATH)/src/github.com/pinpt/schemagen && go run main.go golang $(BASEDIR)/definition $(BASEDIR)/model && \
	[ -d $(GOPATH)/src/github.com/pinpt/pipeline/vendor ] && \
	rm -rf $(GOPATH)/src/github.com/pinpt/pipeline/vendor/github.com/pinpt/integration-sdk/util && \
	rm -rf $(GOPATH)/src/github.com/pinpt/pipeline/vendor/github.com/pinpt/integration-sdk/model && \
	cp -r $(BASEDIR)/model $(GOPATH)/src/github.com/pinpt/pipeline/vendor/github.com/pinpt/integration-sdk && \
	cp -r $(BASEDIR)/util $(GOPATH)/src/github.com/pinpt/pipeline/vendor/github.com/pinpt/integration-sdk || exit 0

dependencies: install-dep
	@rm -rf $(BASEDIR)/.vendor-new $(BASEDIR)/vendor
	@dep ensure

install-dep: install-go
ifeq (, $(shell which dep))
	@brew install dep
endif

install-go:
ifeq (, $(shell which go))
	$(error "No Golang found. Install from https://golang.org/doc/install)
endif
ifneq (1.12,$(firstword $(sort $(go version | awk '{print $3}' | sed 's/go//g') 1.12)))
	$(error "Golang version needs to be 1.12 or greater. Install latest from https://golang.org/doc/install)
endif

.PHONY: help
help:
	@echo -e '\033[0;35mUsage: make <TARGETS>\033[0m'
	@echo ''
	@echo -e '\033[0;32mMain Targets:\033[0m'
	@echo -e '    \033[0;33mdependencies\033[0m       \033[01;34mSetup and install dependencies for the local dev environment.\033[0m'
	@echo -e '    \033[0;33mgenerate\033[0m           \033[01;34mGenerates go code for the schema\033[0m'
	@echo ''
