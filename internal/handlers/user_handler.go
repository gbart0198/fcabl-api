package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gbart/fcabl-api/internal/models"
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// ListUsers handles GET requests to list all users
func (h *Handler) ListUsers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user, err := h.queries.GetUserByEmail(c.Request.Context(), email)
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
		return
	}

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
func (h *Handler) GetUser(c *gin.Context) {
	userID, err := getIntQueryParam("id", c)

	if err != nil {
		if errors.Is(err, ErrParamEmpty) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id parameter is required.",
			})
		} else if errors.Is(err, ErrParamParse) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse user id. Please provide a valid id.",
			})
		}

		return
	}

	user, err := h.queries.GetUserById(c.Request.Context(), userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No user found for id.")
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

func (h *Handler) CreateUser(c *gin.Context) {
	var createUserRequest models.CreateUserRequest // TODO: add password hashing
	var somePassword = "abc123"
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating user.",
		})
		return
	}

	newUser, err := h.queries.CreateUser(c.Request.Context(), createUserRequest.IntoDBModel(somePassword))
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

func (h *Handler) PartialUpdateUser(c *gin.Context) {
	userID, err := getIntPathParam("id", c)
	if err != nil {
		if errors.Is(err, ErrParamEmpty) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Param id is required.",
			})
		} else if errors.Is(err, ErrParamParse) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse id. Please provide a valid id.",
			})
		}
		return
	}
	var updateUserRequest models.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to map update parameters.",
		})
		return
	}
	if updateUserRequest.Email == nil &&
		updateUserRequest.FirstName == nil &&
		updateUserRequest.LastName == nil &&
		updateUserRequest.PhoneNumber == nil &&
		updateUserRequest.Role == nil &&
		updateUserRequest.Password == nil {
		slog.Error("No fields to update were provided.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No fields to update were provided.",
		})
	}

	if err := h.queries.UpdateUser(c.Request.Context(), updateUserRequest.IntoDBModel(userID)); err != nil {
		slog.Error("Failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := getIntPathParam("id", c)
	if err != nil {
		if errors.Is(err, ErrParamEmpty) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Param id is required.",
			})
		} else if errors.Is(err, ErrParamParse) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse id. Please provide a valid id.",
			})
		}
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
