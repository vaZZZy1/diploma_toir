package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App        AppConfig        `mapstructure:"app"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	Postgres   PostgresConfig   `mapstructure:"postgres"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Auth       AuthConfig       `mapstructure:"auth"`
	DataSources []DataSourceConfig `mapstructure:"data_sources"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AuthConfig struct {
	JWTSigningKey    string `mapstructure:"jwt_signing_key"`
	AccessTokenTTL   string `mapstructure:"access_token_ttl"`
	RefreshTokenTTL  string `mapstructure:"refresh_token_ttl"`
}

type DataSourceConfig struct {
	Name   string `mapstructure:"name"`
	Type   string `mapstructure:"type"`
	URL    string `mapstructure:"url"`
	Active bool   `mapstructure:"active"`
}

func LoadConfig(configDir string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configDir)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			defaultConfig := DefaultConfig()
			configBytes, _ := yaml.Marshal(defaultConfig)
			configPath := filepath.Join(configDir, "config.yaml")
			if err := os.WriteFile(configPath, configBytes, 0644); err != nil {
				return nil, fmt.Errorf("не удалось создать файл конфигурации: %w", err)
			}
			return defaultConfig, nil
		}
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("не удалось десериализовать конфигурацию: %w", err)
	}

	if envHost := os.Getenv("POSTGRES_HOST"); envHost != "" {
		config.Postgres.Host = envHost
	}
	if envPort := os.Getenv("POSTGRES_PORT"); envPort != "" {
		config.Postgres.Port = envPort
	}
	if envUser := os.Getenv("POSTGRES_USER"); envUser != "" {
		config.Postgres.Username = envUser
	}
	if envPass := os.Getenv("POSTGRES_PASSWORD"); envPass != "" {
		config.Postgres.Password = envPass
	}
	if envDB := os.Getenv("POSTGRES_DB"); envDB != "" {
		config.Postgres.DBName = envDB
	}
	if envRedisHost := os.Getenv("REDIS_HOST"); envRedisHost != "" {
		config.Redis.Host = envRedisHost
	}
	if envRedisPort := os.Getenv("REDIS_PORT"); envRedisPort != "" {
		config.Redis.Port = envRedisPort
	}

	return &config, nil
}

func DefaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:    "reference-service",
			Version: "1.0.0",
			Env:     "development",
			Port:    8081,
		},
		Logger: LoggerConfig{
			Level: "debug",
		},
		Postgres: PostgresConfig{
			Host:     "postgres",
			Port:     "5432",
			Username: "postgres",
			Password: "postgres",
			DBName:   "aircraft_maintenance",
			SSLMode:  "disable",
		},
		Redis: RedisConfig{
			Host:     "redis",
			Port:     "6379",
			Password: "",
			DB:       0,
		},
		Auth: AuthConfig{
			JWTSigningKey:   "default-secret-key",
			AccessTokenTTL:  "15m",
			RefreshTokenTTL: "24h",
		},
		DataSources: []DataSourceConfig{
			{
				Name:   "Default",
				Type:   "csv",
				URL:    "data/default.csv",
				Active: true,
			},
		},
	}
}