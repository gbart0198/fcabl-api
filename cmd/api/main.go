package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gbart/fcabl-api/internal/auth"
	"github.com/gbart/fcabl-api/internal/config"
	"github.com/gbart/fcabl-api/internal/db"
	"github.com/gbart/fcabl-api/internal/handlers"
	"github.com/gbart/fcabl-api/router"
)

func main() {
	fmt.Println("Initializing application...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Validate required config
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database
	pg, err := db.NewPG(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := pg.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connection established")

	// Initialize JWT service
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTExpirationHours)

	// Initialize handlers
	handler := handlers.NewHandler(pg, jwtService, cfg)

	// Setup router
	r := router.SetupRouter(handler, cfg.FrontendURL, jwtService)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Server listening on %s\n", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
