package handlers

import (
	"net/http"

	"github.com/gbart/fcabl-api/internal/service"
	"github.com/gin-gonic/gin"
)

type GameDetailHandler struct {
	service service.GameDetailService
}

func NewGameDetailHandler(service service.GameDetailService) *GameDetailHandler {
	return &GameDetailHandler{service: service}
}

func (h *GameDetailHandler) ListGameDetails(c *gin.Context) {
	gameDetails, err := h.service.ListGameDetails(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch game details.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gameDetails,
	})
}
