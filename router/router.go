// Package router provides the SetupRouter function, which creates the gin.Engine,
// adds the middleware, sets up the routes, and provides the gin.Engine
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("api/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
