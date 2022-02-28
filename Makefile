#!/usr/bin/env bash
.PHONY: help build

default: help

test: ## run tests
	go test ./... -coverprofile=coverage.out

coverage: ## compute global coverage
	$(MAKE) test
	go tool cover -func=coverage.out | tail -1

live: ## run live app with example as input
		go run cmd/live/main.go -in example.xml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
