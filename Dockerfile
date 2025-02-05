# Dockerfile
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o wav-to-flac-converter main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./wav-to-flac-converter"]
