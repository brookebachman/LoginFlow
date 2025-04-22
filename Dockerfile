# -------- Build Stage --------
    FROM golang:1.24.2 AS builder

    # Set working directory
    WORKDIR /app
    
    # Copy go.mod and go.sum first, for caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the code
    COPY . .
    
    # Cross-compile for Linux (inside container)
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    RUN go build -o main .
    
    # -------- Runtime Stage --------
    FROM alpine:latest
    
    WORKDIR /root/
    
    # Copy binary from builder
    COPY --from=builder /app/main .
    
    # Copy db folder
    COPY --from=builder /app/db ./db
    
    # Expose your port (update if not 8080)
    EXPOSE 8080
    
    # Start the binary
    CMD ["./main"]
    