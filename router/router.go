// Package router provides the SetupRouter function, which creates the gin.Engine,
// adds the middleware, sets up the routes, and provides the gin.Engine
package router

import (
	"net/http"

	"github.com/gbart/fcabl-api/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	// TODO: Setup environment-based cors configuration
	r.Use(cors.Default())

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

	// Team routes
	r.GET("api/team/list", h.ListTeams)
	r.GET("api/team/standings", h.GetTeamStandings)
	r.GET("api/team/stats", h.GetTeamStats)
	r.GET("api/team/players", h.GetTeamWithPlayers)
	r.GET("api/team", h.GetTeam)
	r.GET("api/team/players/list", h.ListTeamsWithPlayers)
	r.POST("api/team", h.CreateTeam)
	r.PUT("api/team", h.UpdateTeam)
	r.DELETE("api/team/:id", h.DeleteTeam)

	// Player routes
	r.GET("api/player/list", h.ListPlayers)
	r.GET("api/player/active", h.ListActivePlayers)
	r.GET("api/player/team", h.ListPlayersByTeam)
	r.GET("api/player/free-agents", h.ListFreeAgents)
	r.GET("api/player/with-user", h.GetPlayerWithUser)
	r.GET("api/player/with-team", h.GetPlayerWithTeam)
	r.GET("api/player/list-with-users", h.ListPlayersWithUsers)
	r.GET("api/player", h.GetPlayer)
	r.POST("api/player", h.CreatePlayer)
	r.PUT("api/player", h.UpdatePlayer)
	r.PATCH("api/player/team", h.UpdatePlayerTeam)
	r.PATCH("api/player/registration", h.UpdatePlayerRegistrationStatus)
	r.DELETE("api/player/:id", h.DeletePlayer)

	// Game routes
	r.GET("api/game/list", h.ListGames)
	r.GET("api/game/upcoming", h.ListUpcomingGames)
	r.GET("api/game/past", h.ListPastGames)
	r.GET("api/game/team", h.ListGamesByTeam)
	r.GET("api/game/with-teams", h.GetGameWithTeams)
	r.GET("api/game/list-with-teams", h.ListGamesWithTeams)
	r.GET("api/game/schedule", h.ListTeamSchedule)
	r.GET("api/game", h.GetGame)
	r.POST("api/game", h.CreateGame)
	r.PUT("api/game", h.UpdateGame)
	r.PUT("api/game/status", h.UpdateGameScoreAndStatus)
	r.PATCH("api/game/time", h.UpdateGameTime)
	r.DELETE("api/game/:id", h.DeleteGame)

	// Payment routes
	r.GET("api/payment/list", h.ListPayments)
	r.GET("api/payment/player", h.ListPaymentsByPlayer)
	r.GET("api/payment/status-filter", h.ListPaymentsByStatus)
	r.GET("api/payment/with-player", h.GetPaymentWithPlayer)
	r.GET("api/payment/list-with-players", h.ListPaymentsWithPlayerInfo)
	r.GET("api/payment/summary", h.GetPlayerPaymentSummary)
	r.GET("api/payment", h.GetPayment)
	r.POST("api/payment", h.CreatePayment)
	r.PATCH("api/payment/status", h.UpdatePaymentStatus)
	r.DELETE("api/payment/:id", h.DeletePayment)

	// TODO: SETUP MIDDLEWARE FOR CORS, etc.

	return r
}
