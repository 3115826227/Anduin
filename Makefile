.PHONY: run build test fmt vet

run:
	go run .

build:
	go build -o bin/anduin .

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...
