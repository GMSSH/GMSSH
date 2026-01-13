SHELL := /bin/bash

# execute "go mod tidy" on all folders that have go.mod file
.PHONY: tidy
tidy:
	go mod tidy

# execute "golangci-lint" to check code style
.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml