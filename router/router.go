// Package router provides the SetupRouter function, which creates the gin.Engine,
// adds the middleware, sets up the routes, and provides the gin.Engine
package router

import (
	"net/http"

	"github.com/gbart/fcabl-api/internal/auth"
	"github.com/gbart/fcabl-api/internal/handlers"
	"github.com/gbart/fcabl-api/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.Handler, frontendURL string, jwtService *auth.JWTService) *gin.Engine {
	r := gin.Default()

	// CORS configuration for HTTP-only cookies
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true, // Required for cookies
		MaxAge:           12 * 3600,
	}))

	// Public routes
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pongerino",
		})
	})

	// Auth routes (public)
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
		authGroup.POST("/verify", h.VerifyToken)
		authGroup.POST("/logout", h.Logout)
		authGroup.POST("/password-reset/request", h.RequestPasswordReset)
		authGroup.POST("/password-reset/confirm", h.ResetPassword)
	}

	// Public game/team routes
	r.GET("/api/team/list", h.ListTeams)
	r.GET("/api/team/standings", h.GetTeamStandings)
	r.GET("/api/team", h.GetTeam)
	r.GET("/api/game/list", h.ListGames)
	r.GET("/api/game/upcoming", h.ListUpcomingGames)
	r.GET("/api/game/past", h.ListPastGames)
	r.GET("/api/game/schedule", h.ListTeamSchedule)
	r.GET("/api/game/list-with-teams", h.ListGamesWithTeams)
	r.GET("/api/game", h.GetGame)

	// Protected routes (require authentication)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		// User routes
		protected.GET("/user", h.GetUser)

		// Admin-only routes
		admin := protected.Group("")
		admin.Use(middleware.AdminMiddleware())
		{
			// User management
			admin.GET("/user/list", h.ListUsers)
			admin.POST("/user", h.CreateUser)
			admin.PUT("/user", h.UpdateUser)
			admin.DELETE("/user/:id", h.DeleteUser)

			// Team management
			admin.GET("/team/stats", h.GetTeamStats)
			admin.GET("/team/players", h.GetTeamWithPlayers)
			admin.GET("/team/players/list", h.ListTeamsWithPlayers)
			admin.POST("/team", h.CreateTeam)
			admin.PUT("/team", h.UpdateTeam)
			admin.DELETE("/team/:id", h.DeleteTeam)

			// Player management
			admin.GET("/player/list", h.ListPlayers)
			admin.GET("/player/active", h.ListActivePlayers)
			admin.GET("/player/team", h.ListPlayersByTeam)
			admin.GET("/player/free-agents", h.ListFreeAgents)
			admin.GET("/player/with-user", h.GetPlayerWithUser)
			admin.GET("/player/with-team", h.GetPlayerWithTeam)
			admin.GET("/player/list-with-users", h.ListPlayersWithUsers)
			admin.GET("/player", h.GetPlayer)
			admin.POST("/player", h.CreatePlayer)
			admin.PUT("/player", h.UpdatePlayer)
			admin.PATCH("/player/team", h.UpdatePlayerTeam)
			admin.PATCH("/player/registration", h.UpdatePlayerRegistrationStatus)
			admin.DELETE("/player/:id", h.DeletePlayer)

			// Game management
			admin.GET("/game/with-teams", h.GetGameWithTeams)
			admin.GET("/game/team", h.ListGamesByTeam)
			admin.POST("/game", h.CreateGame)
			admin.PUT("/game", h.UpdateGame)
			admin.PUT("/game/status", h.UpdateGameScoreAndStatus)
			admin.PATCH("/game/time", h.UpdateGameTime)
			admin.DELETE("/game/:id", h.DeleteGame)

			// Payment management
			admin.GET("/payment/list", h.ListPayments)
			admin.GET("/payment/player", h.ListPaymentsByPlayer)
			admin.GET("/payment/status-filter", h.ListPaymentsByStatus)
			admin.GET("/payment/with-player", h.GetPaymentWithPlayer)
			admin.GET("/payment/list-with-players", h.ListPaymentsWithPlayerInfo)
			admin.GET("/payment/summary", h.GetPlayerPaymentSummary)
			admin.GET("/payment", h.GetPayment)
			admin.POST("/payment", h.CreatePayment)
			admin.PATCH("/payment/status", h.UpdatePaymentStatus)
			admin.DELETE("/payment/:id", h.DeletePayment)
		}
	}

	return r
}
