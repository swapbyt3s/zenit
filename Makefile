SHELL := /bin/bash
.PHONY: help

help: ## Show this help.
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

deps: ## Install dependencies
	go get -u github.com/golang/lint/golint
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/go-ini/ini

lint:
	$(shell go env GOPATH)/bin/golint ./...

tests: ## Run tests
	go test -cover ./...

build: ## Build binary
	go build -ldflags "-s -w" -o zenit main.go

release: ## Create release.
	scripts/release.sh
