# Стадия сборки
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o autoservice-backend .

# Стадия запуска
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/autoservice-backend .
RUN chmod +x ./autoservice-backend
EXPOSE 8080
CMD ["./autoservice-backend"]
