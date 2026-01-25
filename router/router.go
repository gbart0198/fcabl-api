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
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.1.130:5173", "http://192.168.1.21:5173", "http://192.168.1.137:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true, // Required for cookies
		MaxAge:           12 * 3600,
	}))

	public := r.Group("/api")

	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pongerino",
		})
	})

	authGroup := public.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
		authGroup.POST("/verify", h.VerifyToken)
		authGroup.POST("/logout", h.Logout)
		authGroup.POST("/password-reset/request", h.RequestPasswordReset)
		authGroup.POST("/password-reset/confirm", h.ResetPassword)
	}

	publicTeamGroup := public.Group("/team")
	{
		publicTeamGroup.GET("/list", h.ListTeams)
		publicTeamGroup.GET("/", h.GetTeam)
		publicTeamGroup.GET("/players", h.GetTeamWithPlayers)
		publicTeamGroup.GET("/players/list", h.ListTeamsWithPlayers)
	}

	publicGameGroup := public.Group("/game")
	{
		publicGameGroup.GET("/:id", h.GetGame)
		publicGameGroup.GET("/list", h.ListGames)
		publicGameGroup.GET("/schedule", h.ListTeamSchedule)
		publicGameGroup.GET("/schedule/list", h.ListAllSchedules)
		publicGameGroup.GET("/list-with-teams", h.ListGamesWithTeams)
		publicGameGroup.GET("/with-teams", h.GetGameWithTeams)
		publicGameGroup.GET("/team", h.ListGamesByTeam)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		protected.GET("/user/:id", h.GetUserById)
		protected.GET("/user", h.GetUserByEmail)

		protectedUserGroup := protected.Group("/user")
		{
			protectedUserGroup.GET("/:id", h.GetUserById)
			protectedUserGroup.GET("/", h.GetUserByEmail)

		}

		admin := protected.Group("")
		admin.Use(middleware.AdminMiddleware())
		{
			adminUserGroup := admin.Group("/user")
			{
				adminUserGroup.GET("/list", h.ListUsers)
				adminUserGroup.POST("", h.CreateUser)
				adminUserGroup.PUT("", h.UpdateUser)
				adminUserGroup.DELETE("/:id", h.DeleteUser)

			}

			adminTeamGroup := admin.Group("/team")
			{
				adminTeamGroup.POST("/", h.CreateTeam)
				adminTeamGroup.PUT("/", h.UpdateTeam)
				adminTeamGroup.DELETE("/:id", h.DeleteTeam)
			}

			adminPlayerGroup := admin.Group("/player")
			{
				adminPlayerGroup.GET("/list", h.ListPlayers)
				adminPlayerGroup.GET("/team", h.ListPlayersByTeam)
				adminPlayerGroup.GET("/with-user", h.GetPlayerWithUser)
				adminPlayerGroup.GET("/with-team", h.GetPlayerWithTeam)
				adminPlayerGroup.GET("/list-with-users", h.ListPlayersWithUsers)
				adminPlayerGroup.GET("/", h.GetPlayer) // update to use two routes, one with /:id and one with query param for userId
				adminPlayerGroup.POST("/", h.CreatePlayer)
				adminPlayerGroup.PUT("/", h.UpdatePlayer)
				adminPlayerGroup.PATCH("/team", h.UpdatePlayerTeam)
				adminPlayerGroup.PATCH("/registration", h.UpdatePlayerRegistrationStatus)
				adminPlayerGroup.DELETE("/:id", h.DeletePlayer)
			}

			adminGameGroup := admin.Group("/game")
			{
				adminGameGroup.POST("/", h.CreateGame)
				adminGameGroup.PUT("/", h.UpdateGame)
				adminGameGroup.PUT("/status", h.UpdateGameScoreAndStatus)
				adminGameGroup.PATCH("/time", h.UpdateGameTime)
				adminGameGroup.DELETE("/:id", h.DeleteGame)
			}

			adminPaymentGroup := admin.Group("/payment")
			{
				adminPaymentGroup.GET("/list", h.ListPayments)
				adminPaymentGroup.GET("/player", h.ListPaymentsByPlayer)
				adminPaymentGroup.GET("/status-filter", h.ListPaymentsByStatus)
				adminPaymentGroup.GET("/with-player", h.GetPaymentWithPlayer)
				adminPaymentGroup.GET("/list-with-players", h.ListPaymentsWithPlayerInfo)
				adminPaymentGroup.GET("/summary", h.GetPlayerPaymentSummary)
				adminPaymentGroup.GET("/", h.GetPayment)
				adminPaymentGroup.POST("/", h.CreatePayment)
				adminPaymentGroup.PATCH("/status", h.UpdatePaymentStatus)
				adminPaymentGroup.DELETE("/:id", h.DeletePayment)
			}
		}
	}

	return r
}
