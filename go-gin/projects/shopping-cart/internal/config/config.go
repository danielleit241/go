package config

import (
	"fmt"

	"github.com/danielleit241/internal/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type Config struct {
	Environment string
	ServerPort  string
	ApiPrefix   string
	Database    DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Environment: utils.GetEnvOrDefault("ENVIRONMENT", "development"),
		ServerPort:  utils.GetEnvOrDefault("SERVER_PORT", ":8080"),
		ApiPrefix:   utils.GetEnvOrDefault("API_PREFIX", "/api/v1"),
		Database: DatabaseConfig{
			Host:     utils.GetEnvOrDefault("DB_HOST", "localhost"),
			Port:     utils.GetEnvOrDefault("DB_PORT", "5432"),
			User:     utils.GetEnvOrDefault("DB_USER", "postgres"),
			Password: utils.GetEnvOrDefault("DB_PASSWORD", "password"),
			Name:     utils.GetEnvOrDefault("DB_NAME", "mydb"),
			SSLMode:  utils.GetEnvOrDefault("DB_SSL_MODE", "disable"),
		},
	}
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}
