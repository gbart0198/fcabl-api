package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gbart/fcabl-api/internal/auth"
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// RegisterRequest represents the registration payload
type RegisterRequest struct {
	FirstName   string `json:"firstName" binding:"required,min=2"`
	LastName    string `json:"lastName" binding:"required,min=2"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
}

// LoginRequest represents the login payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// VerifyTokenRequest represents token verification payload
type VerifyTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// RequestPasswordResetRequest represents password reset request payload
type RequestPasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordRequest represents password reset payload
type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

// UserResponse represents user data sent to client (no password)
type UserResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Role        string `json:"role"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// Register handles user registration
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("Invalid registration request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}

	// Check if user already exists
	_, err := h.queries.GetUserByEmail(c.Request.Context(), req.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	} else if err != pgx.ErrNoRows {
		slog.Error("Database error checking existing user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Create user
	userID, err := h.queries.CreateUser(c.Request.Context(), repository.CreateUserParams{
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "normal", // Default role
	})
	if err != nil {
		slog.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Fetch created user
	user, err := h.queries.GetUserById(c.Request.Context(), userID)
	if err != nil {
		slog.Error("Failed to fetch created user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Generate JWT token
	token, err := h.jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		slog.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Set HTTP-only cookie
	setAuthCookie(c, token, h.config.JWTExpirationHours)

	// Return response (token is in HTTP-only cookie, not in body)
	userResponse := UserResponse{
		ID:          fmt.Sprintf("%d", user.ID),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339), // Use current time since GetUserById doesn't return UpdatedAt
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": userResponse,
		},
	})
}

// Login handles user authentication
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("Invalid login request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	// Fetch user with password
	user, err := h.queries.GetUserByEmailWithPassword(c.Request.Context(), req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			slog.Error("Database error during login", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		}
		return
	}

	// Verify password
	if err := auth.VerifyPassword(user.PasswordHash, req.Password); err != nil {
		slog.Error("Failed to verify password", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := h.jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		slog.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	// Set HTTP-only cookie
	setAuthCookie(c, token, h.config.JWTExpirationHours)

	// Return response (token is in HTTP-only cookie, not in body)
	userResponse := UserResponse{
		ID:          fmt.Sprintf("%d", user.ID),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Time.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": userResponse,
		},
	})
}

// VerifyToken validates a JWT token and returns user info
func (h *Handler) VerifyToken(c *gin.Context) {
	// Try to get token from cookie first
	token, err := c.Cookie("auth_token")
	if err != nil {
		// If not in cookie, check request body
		var req VerifyTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}
		token = req.Token
	}

	// Validate token
	claims, err := h.jwtService.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Fetch user from database
	user, err := h.queries.GetUserById(c.Request.Context(), claims.UserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			slog.Error("Database error during token verification", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Verification failed"})
		}
		return
	}

	userResponse := UserResponse{
		ID:          fmt.Sprintf("%d", user.ID),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339), // GetUserById doesn't return UpdatedAt
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

// Logout clears the authentication cookie
func (h *Handler) Logout(c *gin.Context) {
	clearAuthCookie(c)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Logged out successfully"}})
}

// RequestPasswordReset initiates the password reset flow
func (h *Handler) RequestPasswordReset(c *gin.Context) {
	var req RequestPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find user by email
	user, err := h.queries.GetUserByEmail(c.Request.Context(), req.Email)
	if err != nil {
		// Don't reveal whether email exists for security
		c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "If the email exists, a reset link has been sent"}})
		return
	}

	// Generate reset token
	resetToken, err := auth.GenerateResetToken()
	if err != nil {
		slog.Error("Failed to generate reset token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	// Store reset token in database
	expiresAt := pgtype.Timestamp{
		Time:  time.Now().Add(time.Minute * time.Duration(h.config.ResetTokenExpirationMin)),
		Valid: true,
	}

	_, err = h.queries.CreatePasswordResetToken(c.Request.Context(), repository.CreatePasswordResetTokenParams{
		UserID:    user.ID,
		Token:     resetToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		slog.Error("Failed to store reset token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	// TODO: Send email with reset link
	// For now, log the token (REMOVE IN PRODUCTION!)
	slog.Info("Password reset requested", "email", req.Email, "token", resetToken)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message":    "If the email exists, a reset link has been sent",
			"resetToken": resetToken, // REMOVE IN PRODUCTION - only for development
		},
	})
}

// ResetPassword completes the password reset flow
func (h *Handler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Verify reset token
	resetToken, err := h.queries.GetPasswordResetToken(c.Request.Context(), req.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	// Hash new password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	// Update user password
	err = h.queries.UpdateUserPassword(c.Request.Context(), repository.UpdateUserPasswordParams{
		PasswordHash: hashedPassword,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ID: resetToken.UserID,
	})
	if err != nil {
		slog.Error("Failed to update password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	// Mark reset token as used
	err = h.queries.MarkPasswordResetTokenAsUsed(c.Request.Context(), req.Token)
	if err != nil {
		slog.Error("Failed to mark reset token as used", "error", err)
		// Continue anyway, password was updated
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Password reset successfully"}})
}

// Helper functions

func setAuthCookie(c *gin.Context, token string, expirationHours int) {
	c.SetCookie(
		"auth_token",         // name
		token,                // value
		expirationHours*3600, // maxAge in seconds
		"/",                  // path
		"",                   // domain (empty = current domain)
		false,                // secure (set to true in production with HTTPS)
		true,                 // httpOnly
	)
}

func clearAuthCookie(c *gin.Context) {
	c.SetCookie(
		"auth_token",
		"",
		-1, // maxAge -1 deletes the cookie
		"/",
		"",
		false,
		true,
	)
}
