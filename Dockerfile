# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o autoservice-backend .

# Stage 2: Create the runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/autoservice-backend .

# Set the binary to be executable
RUN chmod +x ./autoservice-backend

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./autoservice-backend"]
