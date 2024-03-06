package grpc_app

import (
	"fmt"
	"net"
	"strconv"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/config"
	auth_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/server/auth"
	auth_service_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/service/auth"
	auth_postgres_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/storage/postgres/auth"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type App struct {
	log     *logger.Logger
	gRPCSrv *grpc.Server
	port    int
}

func New(log *logger.Logger, cfg *config.Config, postgresDB *sqlx.DB) *App {

	gRPCServer := grpc.NewServer()

	// auth
	authPostgres := auth_postgres_grpc.New(postgresDB)
	authService := auth_service_grpc.New(authPostgres)
	auth_grpc.Register(log, gRPCServer, &authService, cfg.AccessTokenExp, cfg.RefreshTokenExp, cfg.SecretKey)

	return &App{
		log:     log,
		gRPCSrv: gRPCServer,
		port:    cfg.GRPC.Port,
	}
}

func (a *App) MustRun() {

	l, err := net.Listen("tcp", ":"+strconv.Itoa(a.port))

	if err != nil {
		a.log.Ftl(fmt.Sprintf("Failed to listen on port %v, error: %v", ":"+strconv.Itoa(a.port), err.Error()))
	}

	a.log.AppInf(fmt.Sprintf("Server is running on port %v", a.port))

	if err := a.gRPCSrv.Serve(l); err != nil {
		a.log.Ftl("Failed to run the server, error: " + err.Error())
	}
}

func (a *App) GracefulStop(sign string) {
	a.log.AppInf("Graceful stop the server, signal: " + sign)

	a.gRPCSrv.GracefulStop()
}
