// Package postgres implements postgres connection.
package db

import (
	"context"
	"fmt"
	"github.com/golanguzb71/microservice-project-template/config"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

// Postgres -.
type Postgres struct {
	connAttempts int
	connTimeout  time.Duration
	Builder      squirrel.StatementBuilderType
	Db           *pgxpool.Pool
}

func New(cfg *config.Config) (*Postgres, error) {
	response := &Postgres{}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Optional: Customize pool configuration
	config.MaxConns = 1000
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	response.Db = pool
	response.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return response, nil
}
