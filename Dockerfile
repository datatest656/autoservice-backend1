# Строительный этап
FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /autoservice-backend .

# Финальный этап
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /autoservice-backend .
RUN chmod +x ./autoservice-backend
EXPOSE 8080
CMD ["./autoservice-backend"]
