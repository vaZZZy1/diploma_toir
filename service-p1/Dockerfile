FROM golang:1.18-alpine AS builder

WORKDIR /app

# Копирование и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-p1 ./src

# Финальный образ
FROM alpine:3.15

WORKDIR /app

# Копирование исполняемого файла из предыдущего этапа
COPY --from=builder /app/service-p1 .

# Установка timezone
RUN apk --no-cache add tzdata
ENV TZ=Europe/Moscow

EXPOSE 8081

CMD ["./service-p1"] 