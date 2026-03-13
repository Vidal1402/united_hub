package db

import (
  "context"
  "time"

  "github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
  cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
  defer cancel()

  cfg, err := pgxpool.ParseConfig(databaseURL)
  if err != nil {
    return nil, err
  }

  cfg.MaxConns = 20
  cfg.MinConns = 2
  cfg.MaxConnLifetime = 30 * time.Minute
  cfg.MaxConnIdleTime = 5 * time.Minute
  cfg.HealthCheckPeriod = 30 * time.Second

  pool, err := pgxpool.NewWithConfig(cctx, cfg)
  if err != nil {
    return nil, err
  }

  if err := pool.Ping(cctx); err != nil {
    pool.Close()
    return nil, err
  }

  return pool, nil
}