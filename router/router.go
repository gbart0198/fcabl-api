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
		publicTeamGroup.GET("/", h.ListTeams)
		publicTeamGroup.GET("/:id", h.GetTeam)
		publicTeamGroup.GET("/:id/schedule", h.ListTeamSchedule)
		publicTeamGroup.GET("/:id/games", h.ListGamesByTeam)
		publicTeamGroup.GET("/:id/players", h.ListPlayersByTeam)
	}

	publicScheduleGroup := public.Group("/schedule")
	{
		publicScheduleGroup.GET("/", h.ListAllSchedules)
	}

	publicGameGroup := public.Group("/game")
	{
		publicGameGroup.GET("/:id", h.GetGame)
		publicGameGroup.GET("/", h.ListGames)
	}

	publicPlayerGroup := public.Group("/player")
	{
		publicPlayerGroup.GET("/", h.ListPlayers)
		publicPlayerGroup.GET("/:id", h.GetPlayer)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		admin := protected.Group("")
		admin.Use(middleware.AdminMiddleware())
		{
			adminUserGroup := admin.Group("/user")
			{
				adminUserGroup.GET("/:id", h.GetUser)
				adminUserGroup.GET("/", h.ListUsers)
				adminUserGroup.POST("/", h.CreateUser)
				adminUserGroup.PATCH("/:id", h.PartialUpdateUser)
				adminUserGroup.DELETE("/:id", h.DeleteUser)

			}

			adminTeamGroup := admin.Group("/team")
			{
				adminTeamGroup.POST("/", h.CreateTeam)
				adminTeamGroup.PATCH("/:id", h.UpdateTeam)
				adminTeamGroup.DELETE("/:id", h.DeleteTeam)
			}

			adminPlayerGroup := admin.Group("/player")
			{
				adminPlayerGroup.POST("/", h.CreatePlayer)
				adminPlayerGroup.PATCH("/:id", h.UpdatePlayer)
				adminPlayerGroup.DELETE("/:id", h.DeletePlayer)
				adminPlayerGroup.GET("/:id/payment-summary", h.GetPlayerPaymentSummary)
			}

			adminGameGroup := admin.Group("/game")
			{
				adminGameGroup.POST("/", h.CreateGame)
				adminGameGroup.PATCH("/:id", h.UpdateGame)
				adminGameGroup.DELETE("/:id", h.DeleteGame)
			}

			adminPaymentGroup := admin.Group("/payment")
			{
				adminPaymentGroup.GET("/", h.ListPayments)
				adminPaymentGroup.GET("/:id", h.GetPayment)
				adminPaymentGroup.POST("/", h.CreatePayment)
				adminPaymentGroup.PATCH("/:id", h.UpdatePayment)
				adminPaymentGroup.DELETE("/:id", h.DeletePayment)
			}
		}
	}

	return r
}
