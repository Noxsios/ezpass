.DEFAULT_GOAL := build

build:
	go build .

clean:
	rm -f ezpass

words:
	go run gen/main.go > words/gen.go

test:
	go test ./...

.PHONY: build clean words test
