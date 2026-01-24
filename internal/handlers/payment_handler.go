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

// ListPayments handles GET requests to list all payments
func (h *Handler) ListPayments(c *gin.Context) {
	payments, err := h.queries.ListPayments(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payments found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Payment{},
			})
		} else {
			slog.Error("Failed to fetch payments", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch payments",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

// GetPayment handles GET requests for a single payment.
// Query parameters must be used, either id or stripeId.
// Id will take precedence over stripeId.
func (h *Handler) GetPayment(c *gin.Context) {
	paymentIDStr := c.Query("id")
	transactionId := c.Query("transactionId")
	slog.Info("Starting GetPayment", "paymentIdStr", paymentIDStr, "stripeId", transactionId)

	if paymentIDStr == "" {
		if transactionId == "" {
			slog.Warn("Payment ID and Stripe ID are empty.")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide either a payment id or stripe id.",
			})
			return
		}

		payment, err := h.queries.GetPaymentByTransactionId(c.Request.Context(), transactionId)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No payment found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.Payment{},
				})
			} else {
				slog.Error("Error retrieving payment", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving payment.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": payment,
		})
	} else {
		paymentID, err := strconv.ParseInt(paymentIDStr, 10, 64)
		if err != nil {
			slog.Error("Failed to parse payment id", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse payment id. Please provide a valid id.",
			})
			return
		}

		payment, err := h.queries.GetPaymentById(c.Request.Context(), paymentID)
		if err != nil {
			if err == pgx.ErrNoRows {
				slog.Warn("No payment found.")
				c.JSON(http.StatusOK, gin.H{
					"data": []repository.Payment{},
				})
			} else {
				slog.Error("Error retrieving payment", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error retrieving payment.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": payment,
		})
	}
}

// CreatePayment handles POST requests to create a new payment
func (h *Handler) CreatePayment(c *gin.Context) {
	var createPaymentRequest models.CreatePaymentRequest
	if err := c.ShouldBindJSON(&createPaymentRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for creating payment.",
		})
		return
	}

	newPayment, err := h.queries.CreatePayment(c.Request.Context(), createPaymentRequest.IntoDBModel())
	if err != nil {
		slog.Error("Failed to create payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create payment.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newPayment,
	})
}

// UpdatePaymentStatus handles PATCH requests to update a payment's status
func (h *Handler) UpdatePaymentStatus(c *gin.Context) {
	var updatePaymentStatusRequest models.UpdatePaymentStatusRequest
	if err := c.ShouldBindJSON(&updatePaymentStatusRequest); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters for updating payment status.",
		})
		return
	}

	if err := h.queries.UpdatePaymentStatus(c.Request.Context(), updatePaymentStatusRequest.IntoDBModel()); err != nil {
		slog.Error("Failed to update payment status", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update payment status.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeletePayment handles DELETE requests to delete a payment
func (h *Handler) DeletePayment(c *gin.Context) {
	paymentIDStr := c.Param("id")

	paymentID, err := strconv.ParseInt(paymentIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse payment id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse payment id. Please provide a valid id.",
		})
		return
	}

	if err := h.queries.DeletePayment(c.Request.Context(), paymentID); err != nil {
		slog.Error("Failed to delete payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete payment.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// ListPaymentsByPlayer handles GET requests to list payments by player ID
func (h *Handler) ListPaymentsByPlayer(c *gin.Context) {
	playerIDStr := c.Query("playerId")
	slog.Info("Starting ListPaymentsByPlayer", "playerIdStr", playerIDStr)

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

	payments, err := h.queries.ListPaymentsByPlayer(c.Request.Context(), playerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payments found for player.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Payment{},
			})
		} else {
			slog.Error("Failed to fetch payments by player", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch payments by player",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

// ListPaymentsByStatus handles GET requests to list payments by status
func (h *Handler) ListPaymentsByStatus(c *gin.Context) {
	status := c.Query("status")
	slog.Info("Starting ListPaymentsByStatus", "status", status)

	if status == "" {
		slog.Warn("Status is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a status.",
		})
		return
	}

	payments, err := h.queries.ListPaymentsByStatus(c.Request.Context(), status)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payments found with status.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.Payment{},
			})
		} else {
			slog.Error("Failed to fetch payments by status", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch payments by status",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

// GetPaymentWithPlayer handles GET requests for a payment with player details
func (h *Handler) GetPaymentWithPlayer(c *gin.Context) {
	paymentIDStr := c.Query("id")
	slog.Info("Starting GetPaymentWithPlayer", "paymentIdStr", paymentIDStr)

	if paymentIDStr == "" {
		slog.Warn("Payment ID is empty.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a payment id.",
		})
		return
	}

	paymentID, err := strconv.ParseInt(paymentIDStr, 10, 64)
	if err != nil {
		slog.Error("Failed to parse payment id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse payment id. Please provide a valid id.",
		})
		return
	}

	payment, err := h.queries.GetPaymentWithPlayer(c.Request.Context(), paymentID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payment found.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetPaymentWithPlayerRow{},
			})
		} else {
			slog.Error("Error retrieving payment with player", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error retrieving payment with player.",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

// ListPaymentsWithPlayerInfo handles GET requests to list all payments with player info
func (h *Handler) ListPaymentsWithPlayerInfo(c *gin.Context) {
	payments, err := h.queries.ListPaymentsWithPlayerInfo(c.Request.Context())
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payments found.")
			c.JSON(http.StatusOK, gin.H{
				"data": []repository.ListPaymentsWithPlayerInfoRow{},
			})
		} else {
			slog.Error("Failed to fetch payments with player info", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch payments with player info",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

// GetPlayerPaymentSummary handles GET requests to get a player's payment summary
func (h *Handler) GetPlayerPaymentSummary(c *gin.Context) {
	playerIDStr := c.Query("playerId")
	slog.Info("Starting GetPlayerPaymentSummary", "playerIdStr", playerIDStr)

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

	summary, err := h.queries.GetPlayerPaymentSummary(c.Request.Context(), playerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Warn("No payment summary found for player.")
			c.JSON(http.StatusOK, gin.H{
				"data": repository.GetPlayerPaymentSummaryRow{},
			})
		} else {
			slog.Error("Failed to fetch player payment summary", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch player payment summary",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": summary,
	})
}
