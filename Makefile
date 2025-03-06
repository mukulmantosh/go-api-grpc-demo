.PHONY: proto clean build run test

# Go related variables
BINARY_NAME=go-api-grpc-demo
MAIN_FILE=main.go

# Proto related variables
PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/user.proto

# Build the project
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Clean build files
clean:
	@if exist $(BINARY_NAME) del /F $(BINARY_NAME)
	@if exist $(PROTO_DIR)\*.pb.go del /F $(PROTO_DIR)\*.pb.go
	go clean

# Generate proto files
proto:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

# Run the application
run: build
	./$(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Install dependencies
deps:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# All in one command to setup and run
all: clean deps proto build run
