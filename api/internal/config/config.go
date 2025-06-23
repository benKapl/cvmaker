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
	DBURL          string
	Platform       string
	JWTSecret      string
	Port           string
	DevLLMUrl      string
	ProdLLMUrl     string
	DevLLMTimeout  time.Duration
	ProdLLMTimeout time.Duration
	ProdLLMApiKey  string
}

// Load reads configuration from environment variables (.env file is supported)
func Load() (*Config, error) {
	// Load .env file, but don't fail if it doesn't exist (e.g., in production)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or error loading it:", err)
	}

	cfg := Config{
		DevLLMTimeout:  30 * time.Second,
		ProdLLMTimeout: 10 * time.Second,
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
	cfg.DevLLMUrl = os.Getenv("DEV_LLM_URL")
	devLLMTimeoutStr := os.Getenv("DEV_LLM_TIMEOUT_SECONDS")
	if devLLMTimeoutStr != "" {
		seconds, err := strconv.Atoi(devLLMTimeoutStr)
		if err == nil {
			cfg.DevLLMTimeout = time.Duration(seconds) * time.Second
		} else {
			log.Printf("Warning: Invalid LLM_TIMEOUT_SECONDS value '%s', using default %v\n", devLLMTimeoutStr, cfg.DevLLMTimeout)
		}
	}
	cfg.ProdLLMUrl = os.Getenv("PROD_LLM_URL")
	prodLLMTimeoutStr := os.Getenv("PROD_LLM_TIMEOUT_SECONDS")
	if prodLLMTimeoutStr != "" {
		seconds, err := strconv.Atoi(prodLLMTimeoutStr)
		if err == nil {
			cfg.ProdLLMTimeout = time.Duration(seconds) * time.Second
		} else {
			log.Printf("Warning: Invalid PROD_LLM_TIMEOUT_SECONDS value '%s', using default %v\n", prodLLMTimeoutStr, cfg.ProdLLMTimeout)
		}
	}
	cfg.ProdLLMApiKey = os.Getenv("PROD_LLM_API_KEY")

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
	if cfg.DevLLMUrl == "" {
		return nil, errors.New("DEV_LLM_URL must be set")
	}
	if cfg.ProdLLMUrl == "" {
		return nil, errors.New("PROD_LLM_URL must be set")
	}
	if cfg.ProdLLMApiKey == "" {
		return nil, errors.New("PROD_LLM_API_KEY must be set")
	}

	log.Printf("Configuration loaded successfully (Environment: %s)", cfg.Platform)
	return &cfg, nil
}

func GetLLMClient(cfg *Config) llm.LLMClient {
	if cfg.Platform == "dev" {
		return llm.NewOllamaClient(cfg.DevLLMUrl, cfg.DevLLMTimeout)
	}
	return llm.NewMistralClient(cfg.ProdLLMUrl, cfg.ProdLLMApiKey, cfg.ProdLLMTimeout)
}
