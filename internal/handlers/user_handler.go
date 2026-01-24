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

// ListUsers handles GET requests to list all users
func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.queries.ListUsers(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No users found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.User{},
			})
		} else {
			slog.Error("Failed to fetch users", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

// GetUser handles GET requests for a single user.
// Query parameters must be used, either id or email.
// Id will take precedence over email.
func (h *Handler) GetUser(c *gin.Context) {
	userIDStr := c.Query("id")
	email := c.Query("email")
	slog.Info("Starting GetUser", "userIdStr", userIDStr, "email", email)

	if userIDStr == "" {
		if email == "" {
			slog.Warn("User Id and email are empty.")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide either a user id or email.",
			})
			return
		}

		user, err := h.queries.GetUserByEmailWithPassword(c.Request.Context(), email)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No users found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.User{},
				})
			} else {
				slog.Error("Error retreiving user", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retreiving user.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	} else {
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse user id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse user id. Please provide a valid id.",
			})
		}

		user, err := h.queries.GetUserById(c.Request.Context(), userID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No users found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.User{},
				})
			} else {
				slog.Error("Error retreiving user", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving user.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})

	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var createUserRequest models.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating user.",
		})
		return
	}

	newUser, err := h.queries.CreateUser(c.Request.Context(), createUserRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newUser,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	// TODO: Add validation that a user exists
	// ALSO: Consider making the route PUT /api/user/:id
	var updateUserRequest models.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating user.",
		})
		return
	}

	if err := h.queries.UpdateUser(c.Request.Context(), updateUserRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse user id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse user id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeleteUser(c.Request.Context(), userID); err != nil {
		slog.Error("Failed to delete user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
