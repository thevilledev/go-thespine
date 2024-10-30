.DEFAULT_GOAL := all

all: fmt lint test

fmt:
	go fmt $$(go list ./...)

lint: vet
	golangci-lint run

lit: lint

vet:
	go vet $$(go list ./...)

test: test-unit test-fuzz test-bench

test-unit:
	go test -v -race -run ^Test -parallel=8 ./...

test-bench:
	go test -v -benchmem -bench ^Benchmark -parallel=8 ./...

test-fuzz:
	go test -v -race -run ^Fuzz -parallel=8 ./...

.PHONY: fmt lint test