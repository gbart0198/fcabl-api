package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gbart/fcabl-api/internal/models"
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// ListGames handles GET requests to list all games
func (h *Handler) ListGames(c *gin.Context) {
	games, err := h.queries.ListGames(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No games found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Game{},
			})
		} else {
			slog.Error("Failed to fetch games", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch games",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// GetGame handles GET requests for a single game by ID
func (h *Handler) GetGame(c *gin.Context) {
	gameIDStr := c.Query("id")
	slog.Info("Starting GetGame", "gameIdStr", gameIDStr)

	if gameIDStr == "" {
		slog.Warn("Game ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a game id.",
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

	game, err := h.queries.GetGameById(c.Request.Context(), gameID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Game{},
			})
		} else {
			slog.Error("Error retrieving game", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving game.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

// CreateGame handles POST requests to create a new game
func (h *Handler) CreateGame(c *gin.Context) {
	var createGameRequest models.CreateGameRequest
	if err := c.ShouldBindJSON(&createGameRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating game.",
		})
		return
	}

	newGame, err := h.queries.CreateGame(c.Request.Context(), createGameRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create game", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create game.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newGame,
	})
}

// UpdateGame handles PUT requests to update a game
func (h *Handler) UpdateGame(c *gin.Context) {
	var updateGameRequest models.UpdateGameRequest
	if err := c.ShouldBindJSON(&updateGameRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating game.",
		})
		return
	}

	if err := h.queries.UpdateGame(c.Request.Context(), updateGameRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update game", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update game.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeleteGame handles DELETE requests to delete a game
func (h *Handler) DeleteGame(c *gin.Context) {
	gameIDStr := c.Param("id")

	gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse game id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse game id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeleteGame(c.Request.Context(), gameID); err != nil {
		slog.Error("Failed to delete game", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete game.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// ListUpcomingGames handles GET requests to list upcoming games
func (h *Handler) ListUpcomingGames(c *gin.Context) {
	games, err := h.queries.ListUpcomingGames(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No upcoming games found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Game{},
			})
		} else {
			slog.Error("Failed to fetch upcoming games", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch upcoming games",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// ListPastGames handles GET requests to list past games
func (h *Handler) ListPastGames(c *gin.Context) {
	games, err := h.queries.ListPastGames(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No past games found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Game{},
			})
		} else {
			slog.Error("Failed to fetch past games", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch past games",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// ListGamesByTeam handles GET requests to list games by team ID
func (h *Handler) ListGamesByTeam(c *gin.Context) {
	teamIDStr := c.Query("teamId")
	slog.Info("Starting ListGamesByTeam", "teamIdStr", teamIDStr)

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

	games, err := h.queries.ListGamesByTeam(c.Request.Context(), teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No games found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Game{},
			})
		} else {
			slog.Error("Failed to fetch games by team", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch games by team",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// GetGameWithTeams handles GET requests for a game with team details
func (h *Handler) GetGameWithTeams(c *gin.Context) {
	gameIDStr := c.Query("id")
	slog.Info("Starting GetGameWithTeams", "gameIdStr", gameIDStr)

	if gameIDStr == "" {
		slog.Warn("Game ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a game id.",
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

	game, err := h.queries.GetGameWithTeams(c.Request.Context(), gameID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetGameWithTeamsRow{},
			})
		} else {
			slog.Error("Error retrieving game with teams", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving game with teams.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

// ListGamesWithTeams handles GET requests to list all games with team details
func (h *Handler) ListGamesWithTeams(c *gin.Context) {
	games, err := h.queries.ListGamesWithTeams(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No games found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListGamesWithTeamsRow{},
			})
		} else {
			slog.Error("Failed to fetch games with teams", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch games with teams",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// GetGameWithResult handles GET requests for a game with result
func (h *Handler) GetGameWithResult(c *gin.Context) {
	gameIDStr := c.Query("id")
	slog.Info("Starting GetGameWithResult", "gameIdStr", gameIDStr)

	if gameIDStr == "" {
		slog.Warn("Game ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a game id.",
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

	game, err := h.queries.GetGameWithResult(c.Request.Context(), gameID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No game found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetGameWithResultRow{},
			})
		} else {
			slog.Error("Error retrieving game with result", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving game with result.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

// ListGamesWithResults handles GET requests to list all games with results
func (h *Handler) ListGamesWithResults(c *gin.Context) {
	games, err := h.queries.ListGamesWithResults(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No games found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListGamesWithResultsRow{},
			})
		} else {
			slog.Error("Failed to fetch games with results", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch games with results",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

// ListTeamSchedule handles GET requests to list a team's schedule
func (h *Handler) ListTeamSchedule(c *gin.Context) {
	teamIDStr := c.Query("teamId")
	slog.Info("Starting ListTeamSchedule", "teamIdStr", teamIDStr)

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

	schedule, err := h.queries.ListTeamSchedule(c.Request.Context(), teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No schedule found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListTeamScheduleRow{},
			})
		} else {
			slog.Error("Failed to fetch team schedule", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch team schedule",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}

// UpdateGameTime handles PATCH requests to update a game's time
func (h *Handler) UpdateGameTime(c *gin.Context) {
	var updateGameTimeRequest models.UpdateGameTimeRequest
	if err := c.ShouldBindJSON(&updateGameTimeRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating game time.",
		})
		return
	}

	if err := h.queries.UpdateGameTime(c.Request.Context(), updateGameTimeRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update game time", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update game time.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
