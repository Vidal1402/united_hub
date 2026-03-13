package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend_united_hub/internal/cache"
	"backend_united_hub/internal/config"
	"backend_united_hub/internal/db"
	"backend_united_hub/internal/http/router"
	ilog "backend_united_hub/internal/log"
	"backend_united_hub/internal/migrations"
	"github.com/go-playground/validator/v10"
)

func main() {
	logger := ilog.New()

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config load", slog.Any("err", err))
		os.Exit(1)
	}

	if err := os.MkdirAll(cfg.UploadDir, 0o755); err != nil {
		logger.Error("mkdir upload dir", slog.Any("err", err))
		os.Exit(1)
	}

	ctx := context.Background()

	// Mongo
	mongoDB, err := db.NewMongo(ctx, cfg.MongoURI, cfg.MongoDatabase)
	if err != nil {
		logger.Error("mongo connect", slog.Any("err", err))
		os.Exit(1)
	}
	defer mongoDB.Close(context.Background())

	// Redis
	rdb, err := cache.New(ctx, cache.RedisConfig{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	if err != nil {
		logger.Error("redis connect", slog.Any("err", err))
		os.Exit(1)
	}
	defer rdb.Close()

	// Migrations/índices Mongo
	if err := migrations.UpMongo(ctx, mongoDB.Database); err != nil {
		logger.Error("mongo migrations", slog.Any("err", err))
		os.Exit(1)
	}

	v := validator.New()

	h := router.New(router.Deps{
		JWTSecret:      cfg.JWTSecret,
		RequestTimeout: cfg.RequestTimeout,
		DB:             mongoDB.Database,
		Redis:          rdb,
		Validator:      v,
	})

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		logger.Info("server starting", slog.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("server stopped")
}