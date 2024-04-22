# Use the official golang image with the desired version
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o server_central .

# Use a minimal base image to reduce the size of the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage
COPY --from=builder /app/server_central .

# Expose the port on which the server listens
EXPOSE 8080

# Command to run the server binary
CMD ["./server_central"]