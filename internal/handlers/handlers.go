// Package handlers provides HTTP request handlers for various routes.
// Handlers contain the business logic for processing requests and use
// the repository layer to interact with the database.
package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gbart/fcabl-api/internal/db"
	repository "github.com/gbart/fcabl-api/internal/repository/sqlc_gen"
	"github.com/gin-gonic/gin"
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

// HandleListUsers handles GET requests to list all users
func (h *Handler) HandleListUsers(c *gin.Context) {
	users, err := h.queries.ListUsers(c.Request.Context())
	if err != nil {
		slog.Error("Failed to fetch users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
