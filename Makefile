.PHONY: all build vendor test lint

all: lint test

build: vendor
    gb build

vendor:
    gb vendor restore

test: vendor
    gb test -v

lint:
    golangci-lint run src/joke/ src/main/