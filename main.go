package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db        *database.Queries
	platform  string
	JWTsecret string
	JWTissuer string
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
	JWTsecret := os.Getenv("JWT_SECRET")
	if JWTsecret == "" {
		log.Fatal("JWTSecret must be set")
	}
	JWTissuer := os.Getenv("JWT_ISSUER")
	if JWTissuer == "" {
		log.Fatal("JWTissuer must be set")
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
		JWTsecret: JWTsecret,
		JWTissuer: JWTissuer,
		port:      port,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)

	globalMux := LoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: globalMux,
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
