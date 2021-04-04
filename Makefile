.PHONY: clean

all: build

# Build executable for Kale program
build:
	go mod download
	go build --ldflags "-s -w" -o bin/kale ./cmd/kale/main.go

# Build and execute Kale program
start: build
	./bin/kale

# Format Kale source code with Go toolchain
format:
	go fmt ./...

# Clean up binary output folder
clean:
	rm -rf bin/
