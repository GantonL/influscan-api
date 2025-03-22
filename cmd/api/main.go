package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sgl26/influscan-api/internal/database"
	"github.com/sgl26/influscan-api/internal/handlers"
	"github.com/sgl26/influscan-api/internal/middleware"
	"github.com/sgl26/influscan-api/internal/repository"
)

func main() {

	// Initialize Supabase client
	db, err := database.NewSupabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize Supabase client: %v", err)
	}

	// Initialize repositories
	scanRepo := repository.NewScanRepository(db)

	// Initialize handlers
	scanHandler := handlers.NewScanHandler(scanRepo)

	// Create a new server mux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/scans", scanHandler.GetScans)

	// Create server with timeouts and middleware
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.AuthMiddleware(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create shutdown context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
