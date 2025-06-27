.PHONY: help web

help: # me
	@grep -e "^[a-z0-9_-]*:.* #" Makefile | sed -e 's/\(^[a-z0-9_-]*\):.* # \(.*\)/\1: \2/'

web: # Run the web server.
	air

lint: # Run the linter.
	golangci-lint run ./cmd/web
	golangci-lint run ./internal/*

lint-fix: # Run the linter and fix things.
	golangci-lint run --fix ./cmd/web
	golangci-lint run --fix ./internal/*
