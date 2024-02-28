package rest_app

import (
	"context"
	"net/http"
	"time"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/config"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	log     *logger.Logger
	handler *gin.Engine
	server  *http.Server
	port    string
}

func New(log *logger.Logger, cfg *config.Config, db *sqlx.DB) *App {
	router := gin.New()

	return &App{
		log:     log,
		handler: router,
		port:    cfg.REST.Port,
	}
}

func (a *App) MustRun() {
	a.log.AppInf("Rest server started")

	a.server = &http.Server{
		Addr:              ":" + a.port,
		Handler:           a.handler,
		MaxHeaderBytes:    20 << 20,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.log.Ftl("Failed to launch the server: " + err.Error())
	}

}

func (a *App) GracefulStop(sign string) {

	a.log.AppInf("Graceful stop the server, signal: " + sign)

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
