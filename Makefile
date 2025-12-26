.PHONY: run test lint clean

BINARY_NAME := opfetch

run: 
	go run cmd/opfetch/main.go

build: 
	go build -o bin/${BINARY_NAME} cmd/opfetch/main.go

test: 
	go test -v -race ./...

lint: 
	golangci-lint run

clean: 
	go clean
	rm -f bin/${BINARY_NAME}
