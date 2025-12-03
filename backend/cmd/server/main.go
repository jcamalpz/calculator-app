// cmd/server/main.go
package main

import (
	"calculator/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

// corsMiddleware handles CORS headers for local development
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// loggingMiddleware logs all incoming requests
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		log.Printf("Request completed in %v", time.Since(start))
	}
}

func main() {
	h := handlers.NewHandler()

	// API routes
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// Calculator endpoints
	mux.HandleFunc("/api/v1/calculate/add", loggingMiddleware(corsMiddleware(h.Add)))
	mux.HandleFunc("/api/v1/calculate/subtract", loggingMiddleware(corsMiddleware(h.Subtract)))
	mux.HandleFunc("/api/v1/calculate/multiply", loggingMiddleware(corsMiddleware(h.Multiply)))
	mux.HandleFunc("/api/v1/calculate/divide", loggingMiddleware(corsMiddleware(h.Divide)))
	mux.HandleFunc("/api/v1/calculate/power", loggingMiddleware(corsMiddleware(h.Power)))
	mux.HandleFunc("/api/v1/calculate/sqrt", loggingMiddleware(corsMiddleware(h.SquareRoot)))
	mux.HandleFunc("/api/v1/calculate/percentage", loggingMiddleware(corsMiddleware(h.Percentage)))

	port := ":8080"
	fmt.Printf("🚀 Calculator API Server starting on port %s\n", port)
	fmt.Println("📡 Endpoints:")
	fmt.Println("   POST /api/v1/calculate/add")
	fmt.Println("   POST /api/v1/calculate/subtract")
	fmt.Println("   POST /api/v1/calculate/multiply")
	fmt.Println("   POST /api/v1/calculate/divide")
	fmt.Println("   POST /api/v1/calculate/power")
	fmt.Println("   POST /api/v1/calculate/sqrt")
	fmt.Println("   POST /api/v1/calculate/percentage")
	fmt.Println("   GET  /health")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
