package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresPool initializes a parameterized pgxpool.Pool using the provided DB URL.
func NewPostgresPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	// Parse the configuration directly from the database URL
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database url: %w", err)
	}

	// Optimize postgres pool configuration if necessary
	poolConfig.MaxConns = 15

	// Create a new pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}
