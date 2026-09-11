.PHONY: build test

build:
	cd service && go build ./...

test:
	cd service && go test ./...
