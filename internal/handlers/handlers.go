// Package handlers provides HTTP request handlers for various routes.
// Handlers contain the business logic for processing requests and use
// the repository layer to interact with the database.
package handlers

import (
	"github.com/gbart/fcabl-api/internal/db"
	"github.com/gbart/fcabl-api/internal/repository"
)

// Handler holds dependencies for all HTTP handlers
type Handler struct {
	queries *repository.Queries
}

// NewHandler creates a new Handler instance with the provided database connection
func NewHandler(pg *db.Postgres) *Handler {
	return &Handler{
		queries: repository.New(pg.DB),
	}
}
