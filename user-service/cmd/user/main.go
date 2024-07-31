package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"user_service_smart_home/internal/app"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/logger"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	log.Info("Starting service1", slog.Any(
		"config", cfg.RPCPort))
	appv1 := app.NewApp(*cfg, log)

	go appv1.GRPCServer.Run()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("received shutdown signal", slog.String("signal", sig.String()))
	appv1.GRPCServer.Stop()
	log.Info("shutting down server")
}
