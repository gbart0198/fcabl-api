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

// ListGameResults handles GET requests to list all game results
func (h *Handler) ListGameResults(c *gin.Context) {
	results, err := h.queries.ListGameResults(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game results found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.GameResult{},
			})
		} else {
			slog.Error("Failed to fetch game results", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch game results",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

// GetGameResult handles GET requests for a single game result.
// Query parameters must be used, either id or gameId.
// Id will take precedence over gameId.
func (h *Handler) GetGameResult(c *gin.Context) {
	resultIDStr := c.Query("id")
	gameIDStr := c.Query("gameId")
	slog.Info("Starting GetGameResult", "resultIdStr", resultIDStr, "gameIdStr", gameIDStr)

	if resultIDStr == "" {
		if gameIDStr == "" {
			slog.Warn("Result ID and game ID are empty.")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide either a result id or game id.",
			})
			return
		}

		gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse game id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse game id. Please provide a valid id.",
			})
			return
		}

		result, err := h.queries.GetGameResultByGameId(c.Request.Context(), gameID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No game result found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.GameResult{},
				})
			} else {
				slog.Error("Error retrieving game result", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving game result.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	} else {
		resultID, err := strconv.ParseInt(resultIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse result id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse result id. Please provide a valid id.",
			})
			return
		}

		result, err := h.queries.GetGameResultById(c.Request.Context(), resultID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No game result found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.GameResult{},
				})
			} else {
				slog.Error("Error retrieving game result", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving game result.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

// CreateGameResult handles POST requests to create a new game result
func (h *Handler) CreateGameResult(c *gin.Context) {
	var createGameResultRequest models.CreateGameResultRequest
	if err := c.ShouldBindJSON(&createGameResultRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating game result.",
		})
		return
	}

	newResult, err := h.queries.CreateGameResult(c.Request.Context(), createGameResultRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create game result", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create game result.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newResult,
	})
}

// UpdateGameResult handles PUT requests to update a game result
func (h *Handler) UpdateGameResult(c *gin.Context) {
	var updateGameResultRequest models.UpdateGameResultRequest
	if err := c.ShouldBindJSON(&updateGameResultRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating game result.",
		})
		return
	}

	if err := h.queries.UpdateGameResult(c.Request.Context(), updateGameResultRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update game result", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update game result.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeleteGameResult handles DELETE requests to delete a game result
func (h *Handler) DeleteGameResult(c *gin.Context) {
	resultIDStr := c.Param("id")

	resultID, err := strconv.ParseInt(resultIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse result id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse result id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeleteGameResult(c.Request.Context(), resultID); err != nil {
		slog.Error("Failed to delete game result", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete game result.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetGameResultWithTeams handles GET requests for a game result with team details
func (h *Handler) GetGameResultWithTeams(c *gin.Context) {
	resultIDStr := c.Query("id")
	slog.Info("Starting GetGameResultWithTeams", "resultIdStr", resultIDStr)

	if resultIDStr == "" {
		slog.Warn("Result ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a result id.",
		})
		return
	}

	resultID, err := strconv.ParseInt(resultIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse result id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse result id. Please provide a valid id.",
		})
		return
	}

	result, err := h.queries.GetGameResultWithTeams(c.Request.Context(), resultID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game result found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetGameResultWithTeamsRow{},
			})
		} else {
			slog.Error("Error retrieving game result with teams", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving game result with teams.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// ListGameResultsWithTeams handles GET requests to list all game results with team details
func (h *Handler) ListGameResultsWithTeams(c *gin.Context) {
	results, err := h.queries.ListGameResultsWithTeams(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game results found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListGameResultsWithTeamsRow{},
			})
		} else {
			slog.Error("Failed to fetch game results with teams", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch game results with teams",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

// ListTeamGameResults handles GET requests to list game results for a specific team
func (h *Handler) ListTeamGameResults(c *gin.Context) {
	teamIDStr := c.Query("teamId")
	slog.Info("Starting ListTeamGameResults", "teamIdStr", teamIDStr)

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

	results, err := h.queries.ListTeamGameResults(c.Request.Context(), teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game results found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListTeamGameResultsRow{},
			})
		} else {
			slog.Error("Failed to fetch team game results", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch team game results",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

// GetTeamRecord handles GET requests to get a team's record (wins, losses, draws)
func (h *Handler) GetTeamRecord(c *gin.Context) {
	teamIDStr := c.Query("teamId")
	slog.Info("Starting GetTeamRecord", "teamIdStr", teamIDStr)

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

	record, err := h.queries.GetTeamRecord(c.Request.Context(), teamIDPgType)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No record found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetTeamRecordRow{},
			})
		} else {
			slog.Error("Failed to fetch team record", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch team record",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": record,
	})
}
