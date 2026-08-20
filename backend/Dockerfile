# 1. Сборка
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd/main.go

# 2. Финальный образ
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/main /app/main

# Копируем .env (если есть)
COPY .env /app/.env

# Открываем порт
EXPOSE 8080

# Запускаем
CMD ["./main"]