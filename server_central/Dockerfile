# Use an official Golang runtime as the base image
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image to reduce the size of the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage
COPY --from=builder /app/main .

# Expose the port that your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]