.PHONY: build
build:
	go build -v ./cmd/apiserver


.DEFAULT_GOAL := build

.PHONY: test
test:
	go test -v -timeout 30s ./...