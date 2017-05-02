GO ?= go

GIT_TAG = $(shell git describe --tags --abbrev=0 --dirty)
BUILD_OPTS = -ldflags "-s -X main.VERSION=v$(GIT_TAG)"

all: clean build

build:
	$(GO) build -v -i -o leeroy $(BUILD_OPTS)

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -v -i -o leeroy $(BUILD_OPTS)

clean:
	rm -f leeroy

.PHONY: all build build-linux clean
