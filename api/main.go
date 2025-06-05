package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/benKapl/cvmaker_api/internal/config"
	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/handlers"
	"github.com/benKapl/cvmaker_api/internal/llm"
	_ "github.com/lib/pq"
)

func main() {
	// Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration: %v", err)
	}

	// Initialize Dependencies
	dbConn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Error connecting to database: %v", err)
	}
	defer dbConn.Close() // Ensure connection is closed on exit

	if err := dbConn.Ping(); err != nil {
		log.Fatal("Error pinging database: %v", err)
	} else {
		log.Println("Database connection verified")
	}

	dbQueries := database.New(dbConn)

	llmClient := llm.NewClient(cfg.LLMTimeout) // significant timeout to handle llm response time

	api := handlers.NewAPI(dbQueries, llmClient, cfg.JWTSecret, cfg.Platform)

	// prompt := llm.OfferPromptStart + "I am the company Decathlon, i am recruiting a salesman. The salary is 10.000 euros per year. You misson will be to sell our products" + llm.OfferPromptEnd
	// go func() {
	// 	res, err := apiCfg.llmClient.Generate(prompt, llm.OfferFormat)
	// 	if err != nil {
	// 		log.Fatalf("Error: %s", err)
	// 	}
	// 	fmt.Println(res)
	// }()

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	globalMux := handlers.LoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: globalMux,
	}

	log.Printf("Starting server on port %s\n", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
