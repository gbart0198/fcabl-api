package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gbart/fcabl-api/internal/auth"
	"github.com/gbart/fcabl-api/internal/models"
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// GetUsers handles GET requests to list all users
func (h *Handler) GetUsers(c *gin.Context) {
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
	var createUserRequest models.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating user.",
		})
		return
	}

	if err := auth.VerifyPasswordStrength(createUserRequest.Password); err != nil {
		slog.Error("Password failed strength check", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating user.",
		})
		return
	}
	hashedPassword, err := auth.HashPassword(createUserRequest.Password)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user.",
		})
		return
	}
	newUser, err := h.queries.CreateUser(c.Request.Context(), createUserRequest.IntoDBModel(hashedPassword))
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
	var updateUserRequest models.PartialUpdateUserRequest
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

	if updateUserRequest.Email != nil {
		if err := h.queries.UpdateUserEmail(c.Request.Context(), repository.UpdateUserEmailParams{
			Email: *updateUserRequest.Email,
			ID:    userID,
		}); err != nil {
			slog.Error("Failed to update user email", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
	}
	if updateUserRequest.FirstName != nil && updateUserRequest.LastName != nil {
		if err := h.queries.UpdateUserName(c.Request.Context(), repository.UpdateUserNameParams{
			FirstName: *updateUserRequest.FirstName,
			LastName:  *updateUserRequest.LastName,
			ID:        userID,
		}); err != nil {
			slog.Error("Failed to update user name", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
	}
	if updateUserRequest.PhoneNumber != nil {
		if err := h.queries.UpdateUserPhoneNumber(c.Request.Context(), repository.UpdateUserPhoneNumberParams{
			PhoneNumber: *updateUserRequest.PhoneNumber,
			ID:          userID,
		}); err != nil {
			slog.Error("Failed to update user phone number", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
	}
	if updateUserRequest.Role != nil {
		if err := h.queries.UpdateUserRole(c.Request.Context(), repository.UpdateUserRoleParams{
			Role: *updateUserRequest.Role,
			ID:   userID,
		}); err != nil {
			slog.Error("Failed to update user role", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
	}
	if updateUserRequest.Password != nil {
		if err := auth.VerifyPasswordStrength(*updateUserRequest.Password); err != nil {
			slog.Error("Password failed strength check", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameters for updating user.",
			})
			return
		}
		hashedPassword, err := auth.HashPassword(*updateUserRequest.Password)
		if err != nil {
			slog.Error("Failed to hash password", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
		if err := h.queries.UpdateUserPassword(c.Request.Context(), repository.UpdateUserPasswordParams{
			PasswordHash: hashedPassword,
			ID:           userID,
		}); err != nil {
			slog.Error("Failed to update user password", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user.",
			})
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	userID, err := getIntPathParam("id", c)
	if err != nil {
		if errors.Is(err, ErrParamEmpty) {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Param id is required.",
			})
		} else {
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
	if err := h.queries.UpdateUser(c.Request.Context(), repository.UpdateUserParams{
		Email:       updateUserRequest.Email,
		PhoneNumber: updateUserRequest.PhoneNumber,
		FirstName:   updateUserRequest.FirstName,
		LastName:    updateUserRequest.LastName,
		Role:        updateUserRequest.Role,
		ID:          userID,
	}); err != nil {
		slog.Error("Failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user.",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
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

	c.JSON(http.StatusNoContent, gin.H{})
}
