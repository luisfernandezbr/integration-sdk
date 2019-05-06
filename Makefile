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

generate:
	@go run main.go generate $(BASEDIR)/definition $(BASEDIR)/model
