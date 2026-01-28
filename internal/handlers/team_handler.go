package handlers

import (
	"net/http"
	"strings"

	"github.com/gbart/fcabl-api/internal/service"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	service service.TeamService
}

func NewTeamHandler(service service.TeamService) *TeamHandler {
	return &TeamHandler{
		service: service,
	}
}

func (h *TeamHandler) ListTeams(c *gin.Context) {
	includes := strings.Split(c.Query("includes"), ",")
	teams, err := h.service.ListTeams(c.Request.Context(), includes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch teams",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": teams,
	})
}
