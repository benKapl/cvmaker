package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/benKapl/cvmaker_api/internal/config"
	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/llm"
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

	apiCfg := apiConfig{
		db:        dbQueries,
		platform:  platform,
		JWTSecret: JWTSecret,
		port:      port,
		llmClient: llmClient,
	}

	// prompt := llm.OfferPromptStart + "I am the company Decathlon, i am recruiting a salesman. The salary is 10.000 euros per year. You misson will be to sell our products" + llm.OfferPromptEnd
	// go func() {
	// 	res, err := apiCfg.llmClient.Generate(prompt, llm.OfferFormat)
	// 	if err != nil {
	// 		log.Fatalf("Error: %s", err)
	// 	}
	// 	fmt.Println(res)
	// }()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/reset", apiCfg.handlerReset)
	// Auth routes
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)
	// User history
	mux.Handle("POST /api/raw/hobbies", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerRawHobbiesCreate)))
	mux.Handle("POST /api/raw/stacks", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerRawStacksCreate)))
	mux.Handle("POST /api/raw/educations", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerRawEducationsCreate)))
	mux.Handle("POST /api/raw/experiences", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerRawExperiencesCreate)))
	mux.Handle("POST /api/raw/projects", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerRawProjectsCreate)))
	// Offers management
	mux.Handle("POST /api/offers", apiCfg.AuthenticateMiddleware(http.HandlerFunc(apiCfg.handlerOffersCreate)))

	globalMux := LoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: globalMux,
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
