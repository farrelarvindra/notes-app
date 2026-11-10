package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/farrelarvindra/notes-app/handler"
	"github.com/farrelarvindra/notes-app/store"
)

type Config struct {
	Port         string
	DatabaseURL  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func loadConfig() Config {
	// Try loading .env from parent directory (project root) or current directory
	_ = godotenv.Load("../.env")
	_ = godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("DB_URL")
	}

	return Config{
		Port:         port,
		DatabaseURL:  dbURL,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}
}

func main() {
	cfg := loadConfig()

	// Initialize database connection with GORM
	db, err := store.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection error: %v\n", err)
	}

	// Ensure underlying connection is closed on shutdown
	sqlDB, err := db.DB()
	if err == nil {
		defer sqlDB.Close()
	}
	log.Println("Connected to PostgreSQL database successfully via GORM.")

	// Initialize handlers with database dependency
	h := handler.New(db)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      h.Routes(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Channel to listen for shutdown signals
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// Run server in a separate goroutine
	go func() {
		log.Printf("Server is running on port %s (http://localhost:%s)\n", cfg.Port, cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	// Block until a signal is received
	sig := <-shutdownChan
	log.Printf("Received signal: %v. Initiating graceful shutdown...\n", sig)

	// Graceful shutdown context with 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown with error: %v\n", err)
	}

	log.Println("Server gracefully stopped.")
}
