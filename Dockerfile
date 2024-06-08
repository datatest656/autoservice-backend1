# База данных stage
FROM mysql:8.0 AS db
ENV MYSQL_ROOT_PASSWORD=rootpassword
ENV MYSQL_DATABASE=autoservice
ENV MYSQL_USER=autoservice_user
ENV MYSQL_PASSWORD=autoservice_password
COPY ./init.sql /docker-entrypoint-initdb.d/
RUN chmod -R 644 /docker-entrypoint-initdb.d/*

# Сборка stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o autoservice-backend

# Финальная stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/autoservice-backend .
COPY --from=db /docker-entrypoint-initdb.d /docker-entrypoint-initdb.d
EXPOSE 8080

# Команда для запуска базы данных и приложения
CMD ["sh", "-c", "docker-entrypoint.sh mysqld & ./autoservice-backend"]
