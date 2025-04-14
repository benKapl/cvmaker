package main

import (
	"cvmaker_api/internal/database"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	db        *database.Queries
	platform  string
	JWTsecret string
	port      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWTSecret must be set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer dbConn.Close()
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		db:        dbQueries,
		platform:  platform,
		JWTsecret: JWTSecret,
		port:      port,
	}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
