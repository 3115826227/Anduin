.PHONY: run build test fmt vet tidy

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

tidy:
	go mod tidy
