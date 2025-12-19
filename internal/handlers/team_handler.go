package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gbart/fcabl-api/internal/models"
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// ListTeams handles GET requests to list all teams
func (h *Handler) ListTeams(c *gin.Context) {
	teams, err := h.queries.ListTeams(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No teams found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Team{},
			})
		} else {
			slog.Error("Failed to fetch teams", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch teams",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": teams,
	})
}

// GetTeam handles GET requests for a single team by ID
func (h *Handler) GetTeam(c *gin.Context) {
	teamIDStr := c.Query("id")
	slog.Info("Starting GetTeam", "teamIdStr", teamIDStr)

	if teamIDStr == "" {
		slog.Warn("Team ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a team id.",
		})
		return
	}

	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse team id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse team id. Please provide a valid id.",
		})
		return
	}

	team, err := h.queries.GetTeamById(c.Request.Context(), teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No team found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Team{},
			})
		} else {
			slog.Error("Error retrieving team", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving team.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": team,
	})
}

// CreateTeam handles POST requests to create a new team
func (h *Handler) CreateTeam(c *gin.Context) {
	var createTeamRequest models.CreateTeamRequest
	if err := c.ShouldBindJSON(&createTeamRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating team.",
		})
		return
	}

	newTeam, err := h.queries.CreateTeam(c.Request.Context(), createTeamRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create team", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create team.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newTeam,
	})
}

// UpdateTeam handles PUT requests to update a team
func (h *Handler) UpdateTeam(c *gin.Context) {
	var updateTeamRequest models.UpdateTeamRequest
	if err := c.ShouldBindJSON(&updateTeamRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating team.",
		})
		return
	}

	if err := h.queries.UpdateTeam(c.Request.Context(), updateTeamRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update team", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update team.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTeam handles DELETE requests to delete a team
func (h *Handler) DeleteTeam(c *gin.Context) {
	teamIDStr := c.Param("id")

	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse team id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse team id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeleteTeam(c.Request.Context(), teamID); err != nil {
		slog.Error("Failed to delete team", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete team.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetTeamStandings handles GET requests for team standings
func (h *Handler) GetTeamStandings(c *gin.Context) {
	standings, err := h.queries.GetTeamStandings(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No team standings found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.GetTeamStandingsRow{},
			})
		} else {
			slog.Error("Failed to fetch team standings", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch team standings",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": standings,
	})
}

// GetTeamStats handles GET requests for team statistics by ID
func (h *Handler) GetTeamStats(c *gin.Context) {
	teamIDStr := c.Query("id")
	slog.Info("Starting GetTeamStats", "teamIdStr", teamIDStr)

	if teamIDStr == "" {
		slog.Warn("Team ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a team id.",
		})
		return
	}

	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse team id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse team id. Please provide a valid id.",
		})
		return
	}

	stats, err := h.queries.GetTeamStats(c.Request.Context(), teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No team stats found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetTeamStatsRow{},
			})
		} else {
			slog.Error("Error retrieving team stats", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving team stats.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stats,
	})
}

// GetTeamWithPlayers handles GET requests for team with its players
func (h *Handler) GetTeamWithPlayers(c *gin.Context) {
	teamIDStr := c.Query("id")
	slog.Info("Starting GetTeamWithPlayers", "teamIdStr", teamIDStr)

	if teamIDStr == "" {
		slog.Warn("Team ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a team id.",
		})
		return
	}

	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse team id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse team id. Please provide a valid id.",
		})
		return
	}

	// Convert int64 to pgtype.Int8
	var teamIDPgType pgtype.Int8
	teamIDPgType.Scan(teamID)

	players, err := h.queries.GetTeamWithPlayers(c.Request.Context(), teamIDPgType)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No players found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Player{},
			})
		} else {
			slog.Error("Error retrieving team players", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving team players.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}

func (h *Handler) ListTeamsWithPlayers(c *gin.Context) {
	slog.Info("Starting ListTeamsWithPlayers")

	rows, err := h.queries.ListTeamsWithPlayers(c.Request.Context())
	if err != nil {
		slog.Error("Error retrieving teams with players", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error retrieving teams with players.",
		})
		return
	}

	teamMap := make(map[int64]*models.TeamWithPlayers)

	for _, row := range rows {
		// Check if team already exists in map
		if _, exists := teamMap[row.ID]; !exists {
			// Create new team entry
			teamMap[row.ID] = &models.TeamWithPlayers{
				ID:            row.ID,
				Name:          row.Name,
				Wins:          row.Wins,
				Losses:        row.Losses,
				Draws:         row.Draws,
				PointsFor:     row.PointsFor,
				PointsAgainst: row.PointsAgainst,
				CreatedAt:     row.CreatedAt,
				UpdatedAt:     row.UpdatedAt,
				Players:       []models.PlayerSimpleDetails{},
			}
		}

		// Add player to team's player list (only if player data exists)
		if row.JerseyNumber.Valid {
			teamMap[row.ID].Players = append(teamMap[row.ID].Players, models.PlayerSimpleDetails{
				JerseyNumber: row.JerseyNumber,
				FirstName:    row.FirstName.String,
				LastName:     row.LastName.String,
			})
		}
	}

	// Convert map to slice
	teams := make([]models.TeamWithPlayers, 0, len(teamMap))
	for _, team := range teamMap {
		teams = append(teams, *team)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": teams,
	})
}
