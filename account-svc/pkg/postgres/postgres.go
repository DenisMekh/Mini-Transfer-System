package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname, cfg.SSLMode)
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	poolCfg.MaxConns = int32(cfg.MaxOpenConns)
	poolCfg.HealthCheckPeriod = 30 * time.Second
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}
	err = RunMigrations(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	return pool, nil
}

// RunMigrations function for auto apply migrations
func RunMigrations(databaseURL string) error {
	migrationURL := strings.Replace(databaseURL, "postgres://", "pgx5://", 1)
	m, err := migrate.New("file://migrations", migrationURL)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
