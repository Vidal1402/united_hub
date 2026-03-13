package cache

import (
  "context"
  "time"

  "github.com/redis/go-redis/v9"
)

type RedisConfig struct {
  Addr     string
  Password string
  DB       int
}

func New(ctx context.Context, cfg RedisConfig) (*redis.Client, error) {
  rdb := redis.NewClient(&redis.Options{
    Addr:         cfg.Addr,
    Password:     cfg.Password,
    DB:           cfg.DB,
    DialTimeout:  5 * time.Second,
    ReadTimeout:  3 * time.Second,
    WriteTimeout: 3 * time.Second,
    PoolTimeout:  6 * time.Second,
  })

  cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()
  if err := rdb.Ping(cctx).Err(); err != nil {
    return nil, err
  }

  return rdb, nil
}