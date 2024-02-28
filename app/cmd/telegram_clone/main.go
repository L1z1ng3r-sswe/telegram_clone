package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/app"
	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/config"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
)

func main() {
	appConfig := config.MustLoad()

	logger := logger.GetLogger()

	application := app.New(logger, appConfig)

	// running the server
	go func() {
		application.GRPCSrv.MustRun()
	}()

	go func() {
		application.RESTSrv.MustRun()
	}()

	// graceful stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signal := <-stop
	application.GRPCSrv.GracefulStop(string(signal.String()))
	application.RESTSrv.GracefulStop(string(signal.String()))
}
