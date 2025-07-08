package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/benKapl/cvmaker-api/internal/config"
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/handlers"
	"github.com/benKapl/cvmaker-api/internal/services"
	_ "github.com/lib/pq"
)

func main() {
	// Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize Dependencies
	dbConn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConn.Close() // Ensure connection is closed on exit

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	} else {
		log.Println("Database connection verified")
	}

	dbQueries := database.New(dbConn)

	llmClient := config.GetLLMClient(cfg)
	log.Println("LLMClient: ", llmClient.String())

	adminSrv := services.NewAdminService(dbQueries, cfg.Platform)
	authSrv := services.NewAuthService(dbQueries, cfg.JWTSecret)
	offerSrv := services.NewOfferService(dbQueries, llmClient)
	profileSrv := services.NewProfileService(dbQueries)
	resumeSrv := services.NewResumeService(dbQueries, llmClient)

	api := handlers.NewAPI(
		adminSrv,
		authSrv,
		offerSrv,
		profileSrv,
		resumeSrv,
	)

	// Setup router and routes
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	globalMux := handlers.LoggingMiddleware(mux)

	// Start Server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: globalMux,
	}

	log.Printf("Starting server on port %s\n", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
