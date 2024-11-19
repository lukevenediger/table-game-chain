#!/usr/bin/make -f

test:
	@echo "--> Running tests"
	go test -v ./...
