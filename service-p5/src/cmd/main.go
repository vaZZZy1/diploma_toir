package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vazy1/analytics-service/internal/controller"
	"github.com/vazy1/analytics-service/internal/repository"
	"github.com/vazy1/analytics-service/internal/service"
	"github.com/vazy1/analytics-service/pkg/config"
	"github.com/vazy1/analytics-service/pkg/logger"
	"github.com/vazy1/analytics-service/pkg/metrics"
	"github.com/vazy1/analytics-service/pkg/postgres"
	"github.com/vazy1/analytics-service/pkg/redis"
	"github.com/vazy1/analytics-service/pkg/server"
)

// @title Аналитика и отчетность API
// @version 1.0
// @description API сервиса формирования аналитики и отчетности

// @contact.name Поддержка API
// @contact.email support@example.com

// @host localhost:8085
// @BasePath /api/v1

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация логгера
	l := logger.NewLogger(cfg.Logger.Level)
	defer l.Sync()

	// Инициализация подключения к PostgreSQL
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		l.Fatal("Ошибка подключения к PostgreSQL", err)
	}
	defer db.Close()

	// Инициализация Redis
	redisClient, err := redis.NewRedisClient(redis.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		l.Fatal("Ошибка подключения к Redis", err)
	}
	defer redisClient.Close()

	// Инициализация метрик Prometheus
	metrics := metrics.NewMetrics(cfg.App.Name)
	metrics.Register()

	// Инициализация репозиториев
	repos := repository.NewRepositories(db)

	// Инициализация сервисов
	services := service.NewServices(service.Deps{
		Repos:       repos,
		RedisClient: redisClient,
		Metrics:     metrics,
		Config:      cfg,
	})

	// Инициализация обработчиков
	handlers := controller.NewHandler(services, l)

	// Инициализация сервера
	srv := server.NewServer(cfg, handlers.Init())

	// Запуск сервера
	go func() {
		if err := srv.Run(); err != nil {
			l.Fatal("Ошибка запуска сервера", err)
		}
	}()

	l.Info("Сервис аналитики и отчетности запущен")

	// Обработка сигналов для корректного завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	l.Info("Сервис аналитики и отчетности останавливается...")

	// Таймаут на завершение открытых соединений
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Ошибка при остановке сервера", err)
	}

	l.Info("Сервис аналитики и отчетности остановлен")
} 