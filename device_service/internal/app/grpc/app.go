package grpcapp

import (
	alarm_server "device_service/internal/grpc/alarm"
	speaker_server "device_service/internal/grpc/speaker"
	tv_server "device_service/internal/grpc/tv"
	alarm_tv_service_usecase "device_service/internal/usecase/alarm_usecase/alarm"
	user_speaker_service_usecase "device_service/internal/usecase/speaker_usecase/speaker"
	"device_service/internal/usecase/tv_usecase/tv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	GRPCServer *grpc.Server
	logger     *slog.Logger
	Port       string
}

func NewApp(port string, logger *slog.Logger,
	tvService user_tv_service_usecase.TVManagementUseCase,
	alarmService alarm_tv_service_usecase.Alarm,
	speakerService user_speaker_service_usecase.SpeakerManagementUseCase) *App {
	grpcServer := grpc.NewServer()
	tv_server.RegisterTVServer(grpcServer, tvService)
	alarm_server.RegisterAlarmServer(grpcServer, alarmService)
	speaker_server.RegisterSpeakerServer(grpcServer,speakerService)
	reflection.Register(grpcServer)
	return &App{
		GRPCServer: grpcServer,
		logger:     logger,
		Port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"
	log := a.logger.With(
		slog.String("method", op),
		slog.String("port", a.Port))

	l, err := net.Listen("tcp", a.Port)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("starting gRPC server on port", slog.String("port", a.Port))
	err = a.GRPCServer.Serve(l)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
func (a *App) Stop() {
	log := a.logger.With("port", a.Port)
	log.Info("stopping server")
	a.GRPCServer.GracefulStop()
}
