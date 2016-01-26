GO ?= go

all: clean build

build:
	$(GO) build -v -i -o leeroy

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -v -i -o leeroy

clean:
	rm -f leeroy

.PHONY: all build build-linux clean
