BIN_DIR ?= bin

ALL_PACKAGES = ./...
COVERPROFILE = coverage.txt

BUILD_NUMBER := $(TRAVIS_BUILD_NUMBER)
BUILD_DATE = $(shell date "+%Y-%m-%d")
BUILD_TIME = $(shell date "+%H:%M:%S")
BUILD_HASH = $(shell git rev-parse --short HEAD)
BUILD_VERSION ?= $(shell git describe --abbrev=0 --tags)

LDFLAGS += -X github.com/insolar/insolar/version.Version=${BUILD_VERSION}
LDFLAGS += -X github.com/insolar/insolar/version.BuildNumber=${BUILD_NUMBER}
LDFLAGS += -X github.com/insolar/insolar/version.BuildDate=${BUILD_DATE}
LDFLAGS += -X github.com/insolar/insolar/version.BuildTime=${BUILD_TIME}
LDFLAGS += -X github.com/insolar/insolar/version.GitHash=${BUILD_HASH}

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/github-progress-tracker -ldflags "${LDFLAGS}" cmd/*.go

test_with_coverage:
	CGO_ENABLED=1 go test --coverprofile=$(COVERPROFILE) --covermode=atomic $(ALL_PACKAGES)