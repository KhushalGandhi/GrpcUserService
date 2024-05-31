# UserService

This is a gRPC service for managing users.

## Build and Run

To build and run the application, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/UserService.git`
2. Change into the project directory: `cd UserService`
3. Build the application: `go build`
4. Run the application: `./UserService`

## Accessing gRPC Service Endpoints

The gRPC service endpoints can be accessed using a gRPC client. Here's an example using the grpcurl tool:

### List Available Services

To list the available services:

```sh
grpcurl -plaintext localhost:50051 list

Get User by ID
grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserByID


Get Users by IDs
grpcurl -plaintext -d '{"ids": [1, 2, 3]}' localhost:50051 user.UserService/GetUsersByID


Search Users
grpcurl -plaintext -d '{"city": "LA", "phone": 1234567890, "married": true}' localhost:50051 user.UserService/SearchUsers




## Configuration Details

The application does not require any specific configuration. However, ensure that port 50051 is available for the gRPC server to listen on.

