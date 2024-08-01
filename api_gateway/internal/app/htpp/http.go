package htppapp

import (
	"api_gate_way/internal/app/router"
	speaker_use_case "api_gate_way/internal/usecase/speaker"
	"api_gate_way/internal/usecase/tv_use_case"
	"api_gate_way/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, handlerService user_usecase.User,
	tv *tv_use_case.TvUseCaseImpl,
	speaker speaker_use_case.SpeakerUseCaseImpl) *App {
	sever := router.RegisterRouter(handlerService, *tv, speaker)
	return &App{
		Port:   port,
		Server: sever,
		Logger: logger,
	}
}

func (app *App) Start() {

	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.Run(app.Port)
	if err != nil {
		log.Info("Failed to start server", "error", err)
	}
}
