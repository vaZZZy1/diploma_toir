FROM golang:1.19-alpine AS builder

WORKDIR /app

# Копирование go.mod и go.sum
COPY src/go.mod src/go.sum ./
RUN go mod download

# Копирование исходного кода
COPY src ./

# Компиляция приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o planning-service ./cmd/main.go

# Финальный образ
FROM alpine:3.16

WORKDIR /app

# Установка необходимых пакетов
RUN apk --no-cache add ca-certificates tzdata

# Копирование бинарного файла из builder
COPY --from=builder /app/planning-service .

# Копирование конфигурационных файлов
COPY src/configs ./configs

# Установка переменных среды
ENV TZ=Europe/Moscow
ENV APP_ENV=production

# Экспонирование порта
EXPOSE 8082

# Запуск сервиса
CMD ["./planning-service"] 