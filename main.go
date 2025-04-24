package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/llm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db        *database.Queries
	platform  string
	JWTSecret string
	port      string
	llmClient llm.Client
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

	llmClient := llm.NewClient(30 * time.Second) // 30 seconds of timeout to handle llm response time

	apiCfg := apiConfig{
		db:        dbQueries,
		platform:  platform,
		JWTSecret: JWTSecret,
		port:      port,
		llmClient: llmClient,
	}

	prompt := "How many fingers do I have ?"
	go func() {
		res, err := apiCfg.llmClient.Generate(prompt)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		fmt.Println(res)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/reset", apiCfg.handlerReset)
	// Auth routes
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)
	// Business routes
	mux.Handle("POST /api/offers", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerOffersCreate)))

	globalMux := LoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: globalMux,
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
