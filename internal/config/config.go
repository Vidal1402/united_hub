package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	MongoURI       string
	MongoDatabase  string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	JWTSecret      string
	RequestTimeout time.Duration
	UploadDir      string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	c := Config{}
	c.Port = getenv("PORT", "8080")

	// Mongo Atlas URI (ex.: mongodb+srv://user:pass@cluster/... )
	c.MongoURI = getenv("MONGODB_URI", "")
	if c.MongoURI == "" {
		// fallback pra compatibilidade antiga: DATABASE_URL
		c.MongoURI = getenv("DATABASE_URL", "")
	}
	if c.MongoURI == "" {
		return Config{}, fmt.Errorf("MONGODB_URI (ou DATABASE_URL) é obrigatório")
	}

	c.MongoDatabase = getenv("MONGODB_DB", "united_hub")

	c.RedisAddr = getenv("REDIS_ADDR", "localhost:6379")
	c.RedisPassword = os.Getenv("REDIS_PASSWORD")
	c.RedisDB = mustInt(getenv("REDIS_DB", "0"))

	c.JWTSecret = os.Getenv("JWT_SECRET")
	if c.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET é obrigatório")
	}

	c.RequestTimeout = time.Duration(mustInt(getenv("REQUEST_TIMEOUT_MS", "8000"))) * time.Millisecond
	c.UploadDir = getenv("UPLOAD_DIR", "./storage")

	return c, nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func mustInt(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return n
}