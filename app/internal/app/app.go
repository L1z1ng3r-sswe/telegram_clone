package app

import (
	"github.com/L1z1ng3r-sswe/telegram_clone/app/db"
	grpc_app "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/app/grpc"
	rest_app "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/app/rest"
	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/config"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
)

type App struct {
	GRPCSrv *grpc_app.App
	RESTSrv *rest_app.App
}

func New(log *logger.Logger, cfg *config.Config) App {
	postgresDB, err := db.NewPostgres(cfg.PostgresPath)
	if err != nil {
		log.Ftl(err.Error())
	}

	GRPCSrv := grpc_app.New(log, cfg, postgresDB)
	RESTSrv := rest_app.New(log, cfg, postgresDB)

	return App{
		GRPCSrv: GRPCSrv,
		RESTSrv: RESTSrv,
	}
}
