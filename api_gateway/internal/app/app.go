package app

import (
	htppapp "api_gate_way/internal/app/htpp"
	"api_gate_way/internal/config"
	clientgrpcserver "api_gate_way/internal/infastructure/client_grpc_server"
	redis_repository "api_gate_way/internal/infastructure/repository/redis"
	alarm_service "api_gate_way/internal/services/alarm"
	speaker_service "api_gate_way/internal/services/speaker"
	tv_service "api_gate_way/internal/services/tv"
	user_service "api_gate_way/internal/services/user"
	alarm_usecase "api_gate_way/internal/usecase/alarm"
	speaker_use_case "api_gate_way/internal/usecase/speaker"
	"api_gate_way/internal/usecase/tv_use_case"
	"api_gate_way/internal/usecase/user_usecase"
	"log/slog"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	redisDb := redis_repository.NewRedis(*cfg)
	serviceUser := user_usecase.NewUserRepo(redisDb, redisDb, redisDb, redisDb)
	client, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		panic(err)
	}
	userServer := user_service.NewUser(*serviceUser, client, logger)
	user := user_usecase.NewUserUseCase(userServer)

	tvS := tv_service.New(logger, client)
	TV := tv_use_case.NewTvUseCase(tvS)
	spS := speaker_service.New(logger, client)
	SP := speaker_use_case.NewSpeakerUseCase(spS)
	aps := alarm_service.NewAlarm(logger, client)
	AP := alarm_usecase.NewAlarm(aps)
	server := htppapp.NewApp(logger, cfg.HttpPort, user, TV, *SP, *AP)
	return &App{
		HTTPApp: server,
	}
}
