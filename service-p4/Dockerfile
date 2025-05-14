FROM gcc:11 AS builder

WORKDIR /app

# Установка зависимостей
RUN apt-get update && apt-get install -y \
    cmake \
    libboost-all-dev \
    libpqxx-dev \
    libpq-dev \
    libcurl4-openssl-dev \
    libssl-dev \
    nlohmann-json3-dev \
    && rm -rf /var/lib/apt/lists/*

# Копирование исходного кода
COPY src /app/src

# Создание директории для сборки
RUN mkdir -p /app/build

# Сборка приложения
WORKDIR /app/build
RUN cmake ../src && make -j$(nproc)

# Финальный образ
FROM debian:bullseye-slim

# Установка runtime зависимостей
RUN apt-get update && apt-get install -y \
    libboost-system1.74.0 \
    libboost-thread1.74.0 \
    libboost-program-options1.74.0 \
    libpq5 \
    libpqxx-6.4 \
    libcurl4 \
    libssl1.1 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Копирование бинарного файла из builder
COPY --from=builder /app/build/inventory_service .

# Копирование конфигурационных файлов
COPY src/config /app/config

# Экспонирование порта
EXPOSE 8084

# Установка переменных среды
ENV TZ=Europe/Moscow
ENV APP_ENV=production

# Запуск сервиса
CMD ["./inventory_service"] 