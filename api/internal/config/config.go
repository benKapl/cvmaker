package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	Platform   string
	JWTSecret  string
	Port       string
	LLMTimeout time.Duration
}

// Load reads configuration from environment variables (.env file is supported)
func Load() (*Config, error) {
	// Load .env file, but don't fail if it doesn't exist (e.g., in production)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or error loading it:", err)
	}

	cfg := Config{
		LLMTimeout: 120 * time.Second, // Default LLM timeout
	}

	cfg.DBURL = os.Getenv("DB_URL")
	cfg.Platform = os.Getenv("PLATFORM")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.Port = os.Getenv("PORT")
	llmTimeoutStr := os.Getenv("LLM_TIMEOUT_SECONDS")
	if llmTimeoutStr != "" {
		seconds, err := strconv.Atoi(llmTimeoutStr)
		if err == nil {
			cfg.LLMTimeout = time.Duration(seconds) * time.Second
		} else {
			log.Printf("Warning: Invalid LLM_TIMEOUT_SECONDS value '%s', using default %v\n", llmTimeoutStr, cfg.LLMTimeout)
		}
	}

	if cfg.DBURL == "" {
		return nil, errors.New("DB_URL must be set")
	}
	if cfg.Platform == "" {
		return nil, errors.New("PLATFORM must be set")
	}
	if cfg.JWTSecret == "" {
		return nil, errors.New("JWT_SECRET must be set")
	}
	if cfg.Port == "" {
		return nil, errors.New("PORT must be set")
	}

	log.Printf("Configuration loaded successfully (Environment: %s)", cfg.Platform)
	return &cfg, nil
}
