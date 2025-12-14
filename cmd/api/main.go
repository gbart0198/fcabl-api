package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gbart/fcabl-api/internal/db"
	"github.com/gbart/fcabl-api/internal/handlers"
	"github.com/gbart/fcabl-api/router"
)

func main() {
	fmt.Println("Initializing application....")

	connString := os.Getenv("DATABASE_URL")

	pg, err := db.NewPG(context.Background(), connString)
	if err := pg.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to start database pool: %v", err)
	}

	handler := handlers.NewHandler(pg)

	r := router.SetupRouter(handler)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	fmt.Println("Application is listenining on port ")
}
