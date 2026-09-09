.PHONY: build test lint cover install clean

build:
	go build -o gcli ./cmd/gcli

test:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out

cover: test
	go tool cover -html=coverage.out

lint:
	golangci-lint run ./...

install: build
	mv gcli $(HOME)/.local/bin/gcli

clean:
	rm -f gcli coverage.out
