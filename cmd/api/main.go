package main

import (
	"log"

	"github.com/gbart/fcabl-api/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
