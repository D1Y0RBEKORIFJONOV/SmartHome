package main

import (
	"api_gate_way/internal/app"
	"api_gate_way/internal/config"
	"api_gate_way/logger"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	application := app.NewApp(log, cfg)
	forever := make(chan bool)
	go application.HTTPApp.Start()
	<-forever
}
