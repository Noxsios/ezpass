.DEFAULT_GOAL := build

build:
	go build .

clean:
	rm -f ezpass

words:
	go run gen/main.go > words/gen.go

test:
	go test ./...

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./... --fix

.PHONY: build clean words test lint lint-fix
