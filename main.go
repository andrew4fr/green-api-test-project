package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"green-api-test-project/client"
	"green-api-test-project/handlers"
	"green-api-test-project/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	apiURL := getEnv("API_URL", "https://3100.api.green-api.com")
	mediaURL := getEnv("MEDIA_URL", "https://3100.api.green-api.com")
	serverAddr := getEnv("SERVER_ADDR", ":8080")

	httpClient := client.New(apiURL, mediaURL)
	svc := service.NewService(httpClient)
	server := handlers.NewServer(svc)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	r.Mount("/api/v1", handlers.HandlerFromMux(server, r))

	log.Printf("Starting server on %s", serverAddr)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
