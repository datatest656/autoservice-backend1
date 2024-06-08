# Используем официальный образ Golang
FROM golang:1.22-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы в рабочую директорию
COPY . .

# Скачиваем зависимости
RUN go mod download

# Сборка приложения
RUN go build -o /autoservice-backend

# Используем минимальный образ для запуска
FROM alpine:latest

# Копируем скомпилированное приложение из builder stage
COPY --from=builder /autoservice-backend /autoservice-backend

# Устанавливаем права на выполнение
RUN chmod +x /autoservice-backend

# Устанавливаем рабочую директорию
WORKDIR /

# Открываем порт 8080
EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["/autoservice-backend"]
