# Go Gateway Proxy

This project is a Go-based gateway that provides both forward and reverse proxy capabilities with additional features like session management, traffic control, and protocol support for HTTP, TCP, and gRPC. It leverages Go's standard libraries along with third-party frameworks like `Gin` and `grpc-go` for efficient protocol handling and service management.

## Features

### 1. **Connection Handling and Multiplexing**
- Nginx-style connection management where each client request is assigned a unique file descriptor (FD).
- Utilizes event-driven architecture with `epoll` for efficient non-blocking I/O operations.
- Supports connection multiplexing in HTTP/2 by using Stream IDs to handle multiple requests simultaneously within the same connection.

### 2. **Request Identification and Routing**
- Differentiates requests using HTTP request headers like `Host`, `URI`, and custom fields (`X-Forwarded-For`, `X-Requested-With`).
- Supports routing by URI path, request method, and custom headers.
- Implements `location` directives similar to Nginx for URI-based request handling and routing to specific backend services.

### 3. **Load Balancing Algorithms**
- **Round Robin**: Distributes requests sequentially among all backend servers.
- **Least Connections**: Directs traffic to the server with the least number of active connections.
- **IP Hash**: Routes requests from the same IP address to the same backend server, ensuring session consistency.

### 4. **Session Persistence**
- **Cookie-Based Session Persistence**: 
  - When a client first connects, a session cookie (e.g., `srv_id=backend1`) is set.
  - For subsequent requests, the gateway reads the cookie and routes the request to the appropriate server based on the `srv_id`.
- **IP Hash-Based Session Persistence**: Uses the client's IP address to ensure that requests are always routed to the same server.

### 5. **Traffic Control and JWT Authentication**
- Implements JWT-based authentication by verifying the token for each request using middleware in `Gin`.
- Provides a dedicated API endpoint for retrieving and validating JWT tokens.
- Supports traffic control based on client attributes, such as IP address or request headers, ensuring robust access management.

### 6. **Protocol Support and Service Management**
- **HTTP Support**: Uses `Gin` framework to handle HTTP requests and responses efficiently.
- **TCP Support**: Utilizes Go's `net` library for managing low-level TCP connections.
- **gRPC Support**: Uses `grpc-go` library to handle gRPC service requests.
- Service configurations are loaded from a database (e.g., MySQL) into memory during startup using an ORM library like `GORM`.

### 7. **Connection Reuse and Keep-Alive**
- Supports keep-alive connections to reduce overhead from frequent connection establishment.
- Implements connection reuse and pooling to optimize resource utilization for high-concurrency scenarios.

## Architecture

1. **Connection Management**:
    - Utilizes an event loop with `epoll` to monitor file descriptors for incoming events.
    - Each active connection is managed through a connection pool, ensuring efficient use of system resources.

2. **Routing and Load Balancing**:
    - The gateway can route requests to different upstream servers based on URI patterns, HTTP methods, or custom header values.
    - Configurable load balancing strategies ensure optimal traffic distribution across servers.

3. **Session Persistence**:
    - Supports sticky sessions through cookie-based or IP hash mechanisms, ensuring that clients maintain sessions with specific backend servers.

4. **Protocol Handling**:
    - Uses different handlers for HTTP (`Gin`), TCP (`net`), and gRPC (`grpc-go`) protocols.
    - The services are configured to handle traffic for each protocol type independently, allowing the gateway to serve as a multi-protocol proxy.

## Getting Started

### Prerequisites

- Go 1.16 or above
- MySQL (or any other compatible database)
- `Gin` framework
- `grpc-go` library

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/username/go-gateway-proxy.git
    cd go-gateway-proxy
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Start the server:
    ```bash
    go run main.go
    ```

### Configuration

- Modify the `config.yaml` file to set up database connection parameters and service configurations.
- Configure load balancing strategies and session persistence options in the `upstream` section of the configuration file.

### API Endpoints

- **JWT Authentication**: `/api/v1/auth/token`
    - Use this endpoint to retrieve or validate JWT tokens for authentication.
  
- **Service Information**: `/api/v1/service/info`
    - Retrieve information about active services, including HTTP, TCP, and gRPC endpoints.

## Contributing

We welcome contributions to enhance the functionality and features of this gateway proxy. Please feel free to submit pull requests or open issues for any bugs or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
