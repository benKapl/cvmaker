package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/benKapl/cvmaker-api/internal/llm"
	"github.com/joho/godotenv"
)

type Config struct {
	DBURL         string
	Platform      string
	JWTSecret     string
	Port          string
	OllamaUrl     string
	OllamaTimeout time.Duration
}

// Load reads configuration from environment variables (.env file is supported)
func Load() (*Config, error) {
	// Load .env file, but don't fail if it doesn't exist (e.g., in production)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or error loading it:", err)
	}

	cfg := Config{
		OllamaTimeout: 30 * time.Second,
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.New("DB_USER must be set")
	}
	dbPwd := os.Getenv("DB_PASSWORD")
	if dbPwd == "" {
		return nil, errors.New("DB_PASSWORD must be set")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, errors.New("DB_HOST must be set")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, errors.New("DB_PORT must be set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, errors.New("DB_NAME must be set")
	}

	cfg.DBURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPwd, dbHost, dbPort, dbName)
	cfg.Platform = os.Getenv("PLATFORM")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.Port = os.Getenv("PORT")
	cfg.OllamaUrl = os.Getenv("OLLAMA_URL")
	ollamaTimeoutStr := os.Getenv("OLLAMA_TIMEOUT_SECONDS")
	if ollamaTimeoutStr != "" {
		seconds, err := strconv.Atoi(ollamaTimeoutStr)
		if err == nil {
			cfg.OllamaTimeout = time.Duration(seconds) * time.Second
		} else {
			log.Printf("Warning: Invalid LLM_TIMEOUT_SECONDS value '%s', using default %v\n", ollamaTimeoutStr, cfg.OllamaTimeout)
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
	if cfg.OllamaUrl == "" {
		return nil, errors.New("OLLAMA_URL must be set")
	}

	log.Printf("Configuration loaded successfully (Environment: %s)", cfg.Platform)
	return &cfg, nil
}

func GetLLMClient(conf *Config) llm.LLMClient {
	if conf.Platform == "dev" {
		return llm.NewOllamaClient(conf.OllamaUrl, conf.OllamaTimeout)
	}
}
