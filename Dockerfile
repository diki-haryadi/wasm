# Stage 1 - Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary for Linux
RUN go build -o main .

# Stage 2 - Create a minimal final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the public directory for static files
COPY --from=builder /app/public ./public

# Expose the port on which the application will run
EXPOSE 8080

# Run the application
ENTRYPOINT ["./main", "serve"]
