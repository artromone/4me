# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the server and CLI
RUN CGO_ENABLED=0 GOOS=linux go build -o task-server ./cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o task-cli ./cmd/cli/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy built binaries
COPY --from=builder /app/task-server .
COPY --from=builder /app/task-cli .

# Install necessary dependencies
RUN apk --no-cache add postgresql-client

# Expose the server port
EXPOSE 8080

# Default command
CMD ["./task-server"]
