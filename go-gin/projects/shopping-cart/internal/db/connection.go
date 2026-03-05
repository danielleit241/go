package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danielleit241/internal/config"
	"github.com/danielleit241/internal/db/sqlc"
	"github.com/danielleit241/pkg/logger"
	"github.com/danielleit241/pkg/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

var DB sqlc.Querier

func InitDB(config *config.Config) error {
	connectionString := config.GetDatabaseDSN()

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	sqlLogger := logger.NewWithPath("internal/logs/sql.log", "debug", config.Environment)

	cfg.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: &pgx.PgxZerologTracer{
			Logger:         sqlLogger,
			SlowQueryLimit: 500 * time.Millisecond,
		},
		LogLevel: tracelog.LogLevelDebug,
	}

	cfg.MaxConns = 25
	cfg.MinConns = 25
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DBPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create database connection pool: %w", err)
	}

	DB = sqlc.New(DBPool)

	if err := DBPool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return nil
}
