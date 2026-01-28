package handlers

import (
	"net/http"

	"github.com/gbart/fcabl-api/internal/service"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	service service.PlayerService
}

func NewPlayerHandler(service service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) ListPlayers(c *gin.Context) {
	players, err := h.service.ListPlayers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch players",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}
