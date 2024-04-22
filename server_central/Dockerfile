# Use the official golang image with the latest version
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire current directory into the container
COPY . .

# Build the Go application
RUN go build -o server_central .

# Expose the port on which the server listens
EXPOSE 8080

# Command to run the server binary
CMD ["./server_central"]