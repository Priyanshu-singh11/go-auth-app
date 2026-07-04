package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri    string
	MongoDbName string
	JwtSecret   string
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{
		MongoUri:    strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDbName: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JwtSecret:   strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}
	if cfg.MongoDbName == "" {
		return Config{}, fmt.Errorf("missing mongoDB name")
	}
	if cfg.MongoUri == "" {
		return Config{}, fmt.Errorf("missing mongo uri")
	}
	if cfg.JwtSecret == "" {
		return Config{}, fmt.Errorf("missing jwt secret")
	}
	return cfg, nil
}
