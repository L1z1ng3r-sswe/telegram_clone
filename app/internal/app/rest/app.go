package rest_app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/config"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	auth_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/server/auth"
	chat_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/server/chat"
	auth_service_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/service/auth"
	chat_service_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/service/chat"
	auth_postgres_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/storage/postgres/auth"
	chat_postgres_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/storage/postgres/chat"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	log     *logger.Logger
	handler *gin.Engine
	server  *http.Server
	port    string
	hub     chat_rest.Hub
}

func New(log *logger.Logger, cfg *config.Config, postgresDB *sqlx.DB) *App {
	router := gin.New()

	// auth dep
	authRepo := auth_postgres_rest.New(postgresDB)
	authService := auth_service_rest.New(authRepo)
	authServer := auth_rest.New(log, authService, cfg.AccessTokenExp, cfg.RefreshTokenExp, cfg.SecretKey)

	// chat dep
	chatRepo := chat_postgres_rest.New(postgresDB)
	chatService := chat_service_rest.New(chatRepo)
	chatServer := chat_rest.New(log, chatService, cfg.SecretKey)

	// auth router
	auth := router.Group("/auth/")
	// auth.Use(authServer.RequesMW())
	{
		auth.POST("/sign-in", authServer.SignIn)
		auth.POST("/sign-up", authServer.SignUp)
	}

	chat := router.Group("/chat")
	// chat.Use(authServer.RequesMW())
	{
		chat.GET("/new-connection", chatServer.CreateConn)
		chat.POST("/new-bi-chat", authServer.IsAuthMW(), chatServer.CreateBIChat)
		chat.POST("/new-community", authServer.IsAuthMW(), chatServer.CreateCommunity)
		chat.POST("/join-community", authServer.IsAuthMW(), chatServer.JoinCommunity)
		chat.GET("/all-clients", chatServer.GetAllConnections)
		chat.GET("/all-communities", chatServer.GetAllCommunity)
		chat.GET("/all-members", chatServer.GetAllMembersOfCommunity)
		chat.GET("/all-bi-messages/:bi-chat-id", chatServer.GetAllBIMessages)
	}

	return &App{
		log:     log,
		handler: router,
		port:    cfg.REST.Port,
		hub:     chatServer.Hub,
	}
}

func (a *App) MustRun() {

	a.server = &http.Server{
		Addr:              ":" + a.port,
		Handler:           a.handler,
		MaxHeaderBytes:    20 << 20,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go a.hub.Run()

	a.log.AppInf(fmt.Sprintf("Server is running on port %v", a.port))
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.log.Ftl("Failed to launch the server: " + err.Error())
	}

}

func (a *App) GracefulStop(sign string) {

	a.log.AppInf("Graceful stop the server, signal: " + sign)

	a.hub.GracefulStop <- struct{}{}

	if a.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := a.server.Shutdown(ctx); err != nil {
			a.log.Ftl("HTTP server shutdown error: " + err.Error())
		} else {
			a.log.AppInf("HTTP server shutdown completed")
		}
	}

}
