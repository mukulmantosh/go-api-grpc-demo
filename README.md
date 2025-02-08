# Go API and gRPC Demo

This project demonstrates a simple service implementing both REST APIs and gRPC endpoints.

## Features

- REST API endpoints using `gin-gonic`
- gRPC service implementation
- Simple user management functionality
- Concurrent support for both HTTP and gRPC protocols

## Prerequisites

- Go 1.16 or later
- Protocol Buffers compiler (protoc) - [Download here](https://github.com/protocolbuffers/protobuf/releases)
- GNU Make for Windows (you can install it via [chocolatey](https://chocolatey.org/): `choco install make` or use [GnuWin32](http://gnuwin32.sourceforge.net/packages/make.htm))

## Getting Started

1. Make sure protoc is in your PATH. You can verify this by running:
```cmd
protoc --version
```

2. Install all required dependencies:
```cmd
make deps
```

3. Generate protocol buffer code:
```cmd
make proto
```

4. Build and run the server:
```cmd
make run
```

Or, to do everything in one step:
```cmd
make all
```

## Available Make Commands

- `make deps` - Install all required dependencies
- `make proto` - Generate Go code from protocol buffer definitions
- `make build` - Build the application
- `make run` - Build and run the application
- `make test` - Run tests
- `make clean` - Clean up generated files
- `make all` - Clean, install dependencies, generate proto files, build and run

## API Endpoints

### REST API (Default port: 8080)
- GET /api/users - List all users
- GET /api/users/:id - Get user by ID
- POST /api/users - Create new user
- PUT /api/users/:id - Update user
- DELETE /api/users/:id - Delete user

### gRPC (Default port: 50051)
The gRPC service provides the following methods:
- `GetUser` - Get user by ID
- `ListUsers` - List all users
- `CreateUser` - Create new user
- `UpdateUser` - Update existing user
- `DeleteUser` - Delete user by ID

## Interacting with gRPC Service

### Using Postman

1. Open Postman and create a new request
2. Change the request type to "gRPC"
3. Enter the server URL: `localhost:50051`
4. Import the proto file from `proto/user.proto`
5. Select the method you want to test
6. Example messages for each method:

```json
// CreateUser
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 30
}

// GetUser
{
  "id": "user-id-here"
}

// ListUsers
{} // Empty request

// UpdateUser
{
  "id": "user-id-here",
  "name": "John Doe Updated",
  "email": "john.updated@example.com",
  "age": 31
}

// DeleteUser
{
  "id": "user-id-here"
}
```

### Using BloomRPC

[BloomRPC](https://github.com/bloomrpc/bloomrpc) is a great alternative to Postman for gRPC:

1. Download and install BloomRPC
2. Import the proto file (`proto/user.proto`)
3. Set the server address to `localhost:50051`
4. Select a method and modify the request message
5. Click "Play" to send the request

### Using grpcurl

[grpcurl](https://github.com/fullstorydev/grpcurl) is a command-line tool for interacting with gRPC servers:

1. Install grpcurl:
```cmd
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

2. List available services:
```cmd
grpcurl -plaintext localhost:50051 list
```

3. Example commands:
```cmd
# List all users
grpcurl -plaintext localhost:50051 proto.UserService/ListUsers

# Create a user
grpcurl -plaintext -d '{"name": "John Doe", "email": "john@example.com", "age": 30}' localhost:50051 proto.UserService/CreateUser

# Get a user
grpcurl -plaintext -d '{"id": "user-id-here"}' localhost:50051 proto.UserService/GetUser

# Update a user
grpcurl -plaintext -d '{"id": "user-id-here", "name": "John Updated", "email": "john@example.com", "age": 31}' localhost:50051 proto.UserService/UpdateUser

# Delete a user
grpcurl -plaintext -d '{"id": "user-id-here"}' localhost:50051 proto.UserService/DeleteUser
```

### Using gRPCui

[gRPCui](https://github.com/fullstorydev/grpcui) provides a web interface for interacting with gRPC services:

1. Install gRPCui:
```cmd
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

2. Start the UI:
```cmd
grpcui -plaintext localhost:50051
```

3. Open your browser and navigate to the URL shown in the console (usually http://localhost:8080)

Choose the tool that best fits your needs:
- Postman: Great for teams already using Postman, nice UI
- BloomRPC: Dedicated gRPC client with a clean interface
- grpcurl: Command-line tool, great for scripts and automation
- gRPCui: Web interface, good for exploration and testing
