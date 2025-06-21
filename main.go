package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"finances.jordis.golang/api"
	"finances.jordis.golang/infrastructure/dbmodels"
	zapLogger "finances.jordis.golang/infrastructure/logging"
	mysqlmembers "finances.jordis.golang/infrastructure/my-sql/members"
	mysqlmoves "finances.jordis.golang/infrastructure/my-sql/moves"
	jwtService "finances.jordis.golang/services/jwt"
	mail_service "finances.jordis.golang/services/mail"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
func main() {
	ConfigRuntime()
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("App panicked: %v", r)
		}
	}()

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
		return
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       os.Getenv("DATABASE_URL"),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("no db")

		os.Exit(1)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&dbmodels.User{}, &dbmodels.Transaction{}, &dbmodels.Group{})

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtService := &jwtService.Service{
		TokenDuration: 2 * time.Hour,
		Secret:        jwtSecret,
	}

	userRepo := mysqlmembers.NewUsersRepoMySQL(db)
	groupsRepo := mysqlmembers.NewGroupsRepoMySQL(db)
	userGroupsRepo := mysqlmembers.NewUsersGroupRepoMySQL(db)
	transactionsRepo := mysqlmoves.NewTransactionsRepoMySQL(db)

	mailService, err := mail_service.New()
	if err != nil {
		fmt.Print("error initializing mail service")

	}

	loggerService, err := zapLogger.NewDevelopmentZapLogger()
	if err != nil {
		fmt.Print("error initializing logger service")
		os.Exit(1)
		return
	}
	repos := api.Repos{
		UsersRepo:        userRepo,
		GroupsRepo:       groupsRepo,
		UsersGroupsRepo:  userGroupsRepo,
		TransactionsRepo: transactionsRepo,
	}
	services := api.Services{
		JwtService:  jwtService,
		Mailservice: mailService,
		Logger:      loggerService,
	}
	app := &api.App{
		Repos:    repos,
		Services: services,
		DB:       db,
	}

	router := api.Router(app)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
