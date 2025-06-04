package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker_api/internal/auth"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	// Default to 200 OK if WriteHeader is never called.
	return &responseRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		bytesWritten:   0,
	}
}

// Write captures the number of bytes written and calls the original Write.
func (rr *responseRecorder) Write(b []byte) (int, error) {
	n, err := rr.ResponseWriter.Write(b)
	rr.bytesWritten += n

	return n, err
}

// WriteHeader captures the status code and calls the original WriteHeader.
func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode // Record the status code
	rr.ResponseWriter.WriteHeader(statusCode)
}

func (cfg *apiConfig) AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
			return
		}

		userID, err := auth.ValidateJWT(token, cfg.JWTSecret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now() // Start timer

		recorder := newResponseRecorder(w)

		// The recorder satisfies the http.ResponseWriter interface.
		next.ServeHTTP(recorder, r)

		// Calculate duration AFTER the handler finishes
		duration := time.Since(startTime)

		// Log combined information including the recorded status code
		log.Printf(
			"Request: %s %s duration=%s status=%d bytes=%d",
			r.Method,
			r.URL.Path,
			duration,
			recorder.statusCode,
			recorder.bytesWritten,
		)
	})
}
