package app

import (
	grpcapp "device_service/internal/app/grpc"
	"device_service/internal/config"
	mongo_alarmongo_alarm "device_service/internal/infastructure/repository/mongodb/alarm"
	mongo_speaker "device_service/internal/infastructure/repository/mongodb/speaker"
	mongo_tvmongo_tv "device_service/internal/infastructure/repository/mongodb/tv"
	alarm_service "device_service/internal/services/alarm"
	speaker_service "device_service/internal/services/speaker"
	tv_service "device_service/internal/services/tv"
	alarm_tv_service_usecase "device_service/internal/usecase/alarm_usecase/alarm"
	"device_service/internal/usecase/alarm_usecase/alarm_repo"
	user_speaker_service_usecase "device_service/internal/usecase/speaker_usecase/speaker"
	speaker_repo_use_case "device_service/internal/usecase/speaker_usecase/speaker_repo"
	"device_service/internal/usecase/tv_usecase/tv"
	"device_service/internal/usecase/tv_usecase/tv_repo"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	tvDb, err := mongo_tvmongo_tv.NewMongoDB(cfg)
	if err != nil {
		panic(err)
	}
	serviceTV := tv_repo_use_case.NewTVRepoUseCase(tvDb, tvDb, tvDb, tvDb)
	service := tv_service.NewTV(logger, serviceTV)
	grpcServer := user_tv_service_usecase.NewTVManagementUseCase(service)

	alarmDb, err := mongo_alarmongo_alarm.NewMongoDB(cfg)
	if err != nil {
		panic(err)
	}
	alarmTv := alarm_repo.NewAlarmUseCase(alarmDb, alarmDb, alarmDb)
	serviceAlarm := alarm_service.NewAlarm(*alarmTv, logger)
	grpcServerAlarm := alarm_tv_service_usecase.NewAlarmManagementUseCase(serviceAlarm)

	db_speaker, err := mongo_speaker.NewMongoDB(cfg)
	if err != nil {
		panic(err)
	}
	serviceSp := speaker_repo_use_case.NewSpeakerRepoUseCase(db_speaker, db_speaker, db_speaker, db_speaker)
	serviceSP := speaker_service.NewSpeaker(logger, serviceSp)
	grpcSpeaker := user_speaker_service_usecase.NewSpeakerManagementUseCase(serviceSP)

	Server := grpcapp.NewApp(cfg.RPCPort, logger, grpcServer, grpcServerAlarm, grpcSpeaker)
	return &App{
		GRPCServer: Server,
	}
}
