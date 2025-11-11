# -----------------------------
# Build stage
# -----------------------------
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Включаем cgo и собираем бинарь
ENV CGO_ENABLED=1
RUN go build -o main ./cmd/main.go

# -----------------------------
# Runtime stage
# -----------------------------
FROM debian:bookworm-slim

WORKDIR /app

# Копируем бинарь и миграции
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env .env

# Включаем системный резолвер
ENV GODEBUG=netdns=cgo

# Запуск агента напрямую
CMD ["./main"]
