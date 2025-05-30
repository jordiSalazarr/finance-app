package api

import (
	"net/http"
	"time"

	logger "finances.jordis.golang/application/services"
	"finances.jordis.golang/domain/members"
	domainGroups "finances.jordis.golang/domain/members/groups"
	domainUsers "finances.jordis.golang/domain/members/users"
	domainTransaction "finances.jordis.golang/domain/moves/transactions"
	jwtService "finances.jordis.golang/services/jwt"
	mail_service "finances.jordis.golang/services/mail"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Repos struct {
	UsersRepo        domainUsers.UserRepository
	GroupsRepo       domainGroups.GroupRepository
	TransactionsRepo domainTransaction.TransactionsRepository
	UsersGroupsRepo  members.UsersGroupRepository
}
type Services struct {
	JwtService  *jwtService.Service
	Mailservice *mail_service.SMPTService
	Logger      logger.Logger
}

type App struct {
	Repos
	Services
	DB *gorm.DB
}

var limiter = rate.NewLimiter(1, 5)

func rateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		c.Abort()
		return
	}
	c.Next()
}
func Router(app *App) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(rateLimiter, gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://v0-personal-finance-app-xi-fawn.vercel.app"}, // origen de tu frontend

		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	{
		auth := router.Group("/auth")
		auth.POST("/signUp", app.RegisterUserHandler)
		auth.POST("/log-in", app.LoginUserHandler)
		auth.POST("/verify-user", app.VerifyUserHandler)
	}
	{
		v1 := router.Group("/v1")
		v1.Use(ExtractJWTFromRequest(app.JwtService))
		v1.GET("/me", app.TellMyName)
		v1.POST("/groups", app.CreateGroup)
		v1.GET("/user-groups", app.GetUserGroupsHandler)
		v1.POST("/join-group", app.JoinGroup)
		v1.POST("/transactions", app.CreateTransactionHandler)
		v1.PATCH("/transactions/:transactionID", app.ResolveTransaction)
		v1.GET("/transactions", app.GetUserTransactions)
		v1.GET("/transactions/group/:group_id", app.GetGroupTransactions)
		v1.GET("/transactions/pending-to-recieve", app.GetPendingToRecieveTransactions)
		v1.GET("/transactions/pending-to-pay", app.GetPendingToPayTransactions)

	}

	return router

}
