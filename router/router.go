// Package router provides the SetupRouter function, which creates the gin.Engine,
// adds the middleware, sets up the routes, and provides the gin.Engine
package router

import (
	"net/http"

	"github.com/gbart/fcabl-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	// Define API routes
	r.GET("api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pongerino",
		})
	})

	// User routes using handlers
	r.GET("api/user/list", h.ListUsers)
	r.GET("api/user", h.GetUser)
	r.POST("api/user", h.CreateUser)
	r.PUT("api/user", h.UpdateUser)
	r.DELETE("api/user/:id", h.DeleteUser)

	// TODO: SETUP MIDDLEWARE FOR CORS, etc.

	return r
}
