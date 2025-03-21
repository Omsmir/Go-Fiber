# Build stage
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the Go Fiber app
RUN go build -o main .

# Run stage (final image with Alpine for smaller size)
FROM alpine:latest

WORKDIR /root/

# Copy built binary from builder
COPY --from=builder /app/main .

# Give execution permission (important!)
RUN chmod +x main

# Expose the port Fiber is running on
EXPOSE 8080

# Run the binary
CMD ["./main"]
