package handlers

import (
	"net/http"
	"strings"

	"github.com/gbart/fcabl-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	includes := strings.Split(c.Query("includes"), ",")
	users, err := h.service.ListUsers(c.Request.Context(), includes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
