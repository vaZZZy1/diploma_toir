FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-gateway ./src

# Финальный образ
FROM alpine:3.15

WORKDIR /app
COPY --from=builder /app/api-gateway .

RUN apk --no-cache add tzdata
ENV TZ=Europe/Moscow

EXPOSE 8080

CMD ["./api-gateway"] 