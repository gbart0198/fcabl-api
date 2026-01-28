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

// ListPlayers handles GET requests to list all players
func (h *Handler) ListPlayers(c *gin.Context) {
	players, err := h.queries.ListPlayers(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No players found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Player{},
			})
		} else {
			slog.Error("Failed to fetch players", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch players",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}

// GetPlayer handles GET requests for a single player.
// Query parameters must be used, either id or userId.
// Id will take precedence over userId.
// update to account for all includes query parameters
func (h *Handler) GetPlayer(c *gin.Context) {
	// TODO: Update query object to pull back user information like name, email, etc.
	playerIDStr := c.Query("id")
	userIDStr := c.Query("userId")
	slog.Info("Starting GetPlayer", "playerIdStr", playerIDStr, "userIdStr", userIDStr)

	if playerIDStr == "" {
		if userIDStr == "" {
			slog.Warn("Player ID and user ID are empty.")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide either a player id or user id.",
			})
			return
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse user id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse user id. Please provide a valid id.",
			})
			return
		}

		player, err := h.queries.GetPlayerByUserId(c.Request.Context(), userID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No player found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.Player{},
				})
			} else {
				slog.Error("Error retrieving player", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving player.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": player,
		})
	} else {
		playerID, err := strconv.ParseInt(playerIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse player id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse player id. Please provide a valid id.",
			})
			return
		}

		player, err := h.queries.GetPlayerById(c.Request.Context(), playerID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No player found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.Player{},
				})
			} else {
				slog.Error("Error retrieving player", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving player.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": player,
		})
	}
}

// CreatePlayer handles POST requests to create a new player
func (h *Handler) CreatePlayer(c *gin.Context) {
	var createPlayerRequest models.CreatePlayerRequest
	if err := c.ShouldBindJSON(&createPlayerRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating player.",
		})
		return
	}

	newPlayer, err := h.queries.CreatePlayer(c.Request.Context(), createPlayerRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create player", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create player.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newPlayer,
	})
}

// UpdatePlayer handles PUT requests to update a player
func (h *Handler) UpdatePlayer(c *gin.Context) {
	var updatePlayerRequest models.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&updatePlayerRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating player.",
		})
		return
	}

	if err := h.queries.UpdatePlayer(c.Request.Context(), updatePlayerRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update player", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update player.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeletePlayer handles DELETE requests to delete a player
func (h *Handler) DeletePlayer(c *gin.Context) {
	playerIDStr := c.Param("id")

	playerID, err := strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse player id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse player id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeletePlayer(c.Request.Context(), playerID); err != nil {
		slog.Error("Failed to delete player", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete player.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// ListPlayersByTeam handles GET requests to list players by team ID
func (h *Handler) ListPlayersByTeam(c *gin.Context) {
	teamIDStr := c.Query("teamId") // update to pull from route params
	slog.Info("Starting ListPlayersByTeam", "teamIdStr", teamIDStr)

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

	players, err := h.queries.ListPlayersByTeam(c.Request.Context(), teamIDPgType)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No players found for team.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Player{},
			})
		} else {
			slog.Error("Failed to fetch players by team", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch players by team",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}

// GetPlayerWithUser handles GET requests for a player with user details
func (h *Handler) GetPlayerWithUser(c *gin.Context) {
	playerIDStr := c.Query("id")
	slog.Info("Starting GetPlayerWithUser", "playerIdStr", playerIDStr)

	if playerIDStr == "" {
		slog.Warn("Player ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a player id.",
		})
		return
	}

	playerID, err := strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse player id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse player id. Please provide a valid id.",
		})
		return
	}

	player, err := h.queries.GetPlayerWithUserInfo(c.Request.Context(), playerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No player found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetPlayerWithUserInfoRow{},
			})
		} else {
			slog.Error("Error retrieving player with user", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving player with user.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": player,
	})
}

// GetPlayerWithTeam handles GET requests for a player with team details
func (h *Handler) GetPlayerWithTeam(c *gin.Context) {
	playerIDStr := c.Query("id")
	slog.Info("Starting GetPlayerWithTeam", "playerIdStr", playerIDStr)

	if playerIDStr == "" {
		slog.Warn("Player ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a player id.",
		})
		return
	}

	playerID, err := strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse player id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse player id. Please provide a valid id.",
		})
		return
	}

	player, err := h.queries.GetPlayerWithUserAndTeamInfo(c.Request.Context(), playerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No player found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetPlayerWithUserAndTeamInfoRow{},
			})
		} else {
			slog.Error("Error retrieving player with team", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving player with team.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": player,
	})
}

// ListPlayersWithUsers handles GET requests to list all players with user details
func (h *Handler) ListPlayersWithUsers(c *gin.Context) {
	players, err := h.queries.ListPlayersWithUsers(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No players found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListPlayersWithUsersRow{},
			})
		} else {
			slog.Error("Failed to fetch players with users", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch players with users",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}

// UpdatePlayerTeam handles PATCH requests to update a player's team
func (h *Handler) UpdatePlayerTeam(c *gin.Context) {
	var updatePlayerTeamRequest models.UpdatePlayerTeamRequest
	if err := c.ShouldBindJSON(&updatePlayerTeamRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating player team.",
		})
		return
	}

	if err := h.queries.UpdatePlayerTeam(c.Request.Context(), updatePlayerTeamRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update player team", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update player team.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePlayerRegistrationStatus handles PATCH requests to update a player's registration status
func (h *Handler) UpdatePlayerRegistrationStatus(c *gin.Context) {
	var updatePlayerRegistrationFeeRequest models.UpdatePlayerRegistrationFeeRequest
	if err := c.ShouldBindJSON(&updatePlayerRegistrationFeeRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating player registration status.",
		})
		return
	}

	if err := h.queries.UpdatePlayerRegistrationFee(c.Request.Context(), updatePlayerRegistrationFeeRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update player registration status", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update player registration status.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
