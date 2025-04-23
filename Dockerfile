# -------- Build Stage --------
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum first, for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build with CGO enabled
ENV CGO_ENABLED=1
RUN go build -o main .

# -------- Runtime Stage --------
FROM alpine:latest

WORKDIR /root/

# Install runtime dependencies
RUN apk add --no-cache sqlite

# Copy binary from builder
COPY --from=builder /app/main .

# Copy db folder
COPY --from=builder /app/db ./db

# Expose your port (update if not 8080)
EXPOSE 8080

# Start the binary
CMD ["./main"]
    