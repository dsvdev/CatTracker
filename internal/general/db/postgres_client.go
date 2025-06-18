package db

import (
	"CatTracker/internal/general/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PostgresClient struct {
	Pool *pgxpool.Pool
}

func NewPostgresClient(ctx context.Context, dbConfig *config.DbConfig) (*PostgresClient, error) {
	cfg, err := pgxpool.ParseConfig(dbConfig.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	// Настройка логов, таймаутов и т.д.
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	return &PostgresClient{Pool: pool}, nil
}

func (p *PostgresClient) Close() {
	p.Pool.Close()
}

func (p *PostgresClient) Ping(ctx context.Context) error {
	return p.Pool.Ping(ctx)
}
