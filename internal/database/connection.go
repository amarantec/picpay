package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func OpenConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	Conn, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}
	return Conn, err
}
