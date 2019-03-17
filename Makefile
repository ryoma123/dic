DIR ?= ./...

default: build

build:
	go build ./cmd/dic

lint:
	golint -set_exit_status $(DIR)
	go vet $(DIR)

test:
	go test $(DIR)

deps:
	go get golang.org/x/lint/golint
	go get github.com/goreleaser/goreleaser

ci: deps test lint
	git diff go.mod

.PHONY: default build lint test deps ci
