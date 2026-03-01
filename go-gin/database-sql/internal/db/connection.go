package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/danielleit241/internal/config"
)

var pool *sql.DB

func Init(config *config.Config) (*sql.DB, error) {
	var err error
	pool, err = sql.Open("postgres", config.GetDatabaseDSN())
	if err != nil {
		return nil, err
	}

	configurePool()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = pool.PingContext(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func configurePool() {
	pool.SetMaxOpenConns(3)
	pool.SetMaxIdleConns(3)
	pool.SetConnMaxLifetime(30 * time.Minute)
	pool.SetConnMaxIdleTime(5 * time.Minute)
}
