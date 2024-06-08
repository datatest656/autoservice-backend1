# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o autoservice-backend .

# Stage 2: Create the final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/autoservice-backend .
EXPOSE 8080
CMD ["./autoservice-backend"]
