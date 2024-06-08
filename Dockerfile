# Используйте официальное изображение Go для сборки
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /autoservice-backend .

# Используйте минимальное изображение Alpine для финального контейнера
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /autoservice-backend .
RUN chmod +x ./autoservice-backend
CMD ["./autoservice-backend"]
