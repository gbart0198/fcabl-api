package handlers

import (
	"net/http"
	"strings"

	"github.com/gbart/fcabl-api/internal/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service service.GameService
}

func NewGameHandler(service service.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) ListGames(c *gin.Context) {
	includes := strings.Split(c.Query("include"), ",")
	games, err := h.service.ListGames(c.Request.Context(), includes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}
