BIN_DIR ?= bin

ALL_PACKAGES = ./...
COVERPROFILE = coverage.txt

BUILD_NUMBER := $(TRAVIS_BUILD_NUMBER)
BUILD_DATE = $(shell date "+%Y-%m-%d")
BUILD_TIME = $(shell date "+%H:%M:%S")
BUILD_HASH = $(shell git rev-parse --short HEAD)
BUILD_VERSION ?= $(shell git describe --abbrev=0 --tags)

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/github-progress-tracker cmd/*.go

test_with_coverage:
	CGO_ENABLED=1 go test --coverprofile=$(COVERPROFILE) --covermode=atomic $(ALL_PACKAGES)