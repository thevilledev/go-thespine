.DEFAULT_GOAL := all

all: fmt lint test

fmt:
	go fmt $$(go list ./...)

lint: vet
	golangci-lint run

lit: lint

vet:
	go vet $$(go list ./...)

test:
	go test -v -race -run ^Test -parallel=8 ./...

.PHONY: fmt lint test