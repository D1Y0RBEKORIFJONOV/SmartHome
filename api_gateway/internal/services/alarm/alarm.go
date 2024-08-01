package alarm_service

import (
	"api_gate_way/internal/config"
	"api_gate_way/internal/entity"
	clientgrpcserver "api_gate_way/internal/infastructure/client_grpc_server"
	"api_gate_way/internal/infastructure/producer"
	"context"
	"encoding/json"
	"errors"
	alarm1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/smart_alarm"
	"log"
	"log/slog"
)

type Alarm struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
	p      *producer.Producer
}

func NewAlarm(logger *slog.Logger, client clientgrpcserver.ServiceClient) *Alarm {
	cfg := config.New()
	p, err := producer.NewProducer(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &Alarm{
		logger: logger,
		client: client,
		p:      p,
	}
}

//type AlarmUseCase interface {
//	AddSmartAlarm(ctx context.Context, alarm entity.AddSmartAlarmReq) (entity.StatusMessage, error)
//	OpenCurtain(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error)
//	CreateAlarmClock(ctx context.Context, req *entity.CreateAlarmClockReq) (entity.StatusMessage, error)
//	OpenDoor(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error)
//	GetRemainingTime(ctx context.Context, userID, deviceName string) (entity.GetUserChannelRes, error)
//}

func (a *Alarm) AddSmartAlarm(ctx context.Context, alarm entity.AddSmartAlarmReq) (entity.StatusMessage, error) {
	const op = "alarm.AddSmartAlarm"
	log := a.logger.With(
		op, "AddSmartAlarm")

	req := entity.AddSmartAlarmReqPub{
		UserID:     alarm.UserID,
		ModelName:  alarm.ModelName,
		DeviceName: alarm.DeviceName,
		MethodName: "Alarm.AddSmartAlarm",
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Pub")
	err = a.p.Pub(ctx, reqBytes, "devices")
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Alarm")

	return entity.StatusMessage{
		Message: "Add smart alarm processing",
	}, nil
}

func (a *Alarm) CreateAlarmClock(ctx context.Context, req *entity.CreateAlarmClockReq) (entity.StatusMessage, error) {
	const op = "alarm.CreateAlarmClock"
	log := a.logger.With(
		op, "CreateAlarmClock")
	log.Info("Called Pub")
	reqA := entity.CreateAlarmClockReqPub{
		UserID:     req.UserID,
		ClockTime:  req.ClockTime,
		DeviceName: req.DeviceName,
		MethodName: "Alarm.CreateAlarmClock",
	}
	reqBytes, err := json.Marshal(reqA)
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Pub")
	err = a.p.Pub(ctx, reqBytes, "devices")
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Alarm")

	return entity.StatusMessage{
		Message: "Add smart alarm processing",
	}, nil
}

func (a *Alarm) OpenDoor(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error) {
	const op = "alarm.OpenDoor"
	log := a.logger.With(
		op, "OpenDoor")
	open := true
	if boolStr == "false" {
		open = false
	}
	message, err := a.client.SmartAlarm().OpenAndClose(ctx, &alarm1.OpenAndCloseReq{
		UserId:     userID,
		DeviceName: deviceName,
		Open:       open,
	})
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Alarm")

	return entity.StatusMessage{
		Message: message.Message,
	}, nil
}

func (a *Alarm) GetRemainingTime(ctx context.Context, userID, deviceName string) (*entity.RemainingTimRes, error) {
	const op = "alarm.GetRemainingTime"
	log := a.logger.With(
		op, "GetRemainingTime")

	resulr, err := a.client.SmartAlarm().GetRemainingTime(ctx, &alarm1.RemainingTimeReq{
		UserId:     userID,
		DeviceName: deviceName,
	})
	if err != nil {
		return nil, errors.New(op + err.Error())
	}
	log.Info("Called Alarm")
	var Result entity.RemainingTimRes
	for _, d := range resulr.Alarms {
		Result.Alarms = append(Result.Alarms, entity.Alarm{
			AlarmTime:     d.AlarmTime,
			RemainingTime: d.RemainingTime,
		})
	}
	Result.Count = resulr.Count

	log.Info("Result>>>>>>>>>>>>>>>>>>", Result)
	log.Info("Result>>>>>>>>>>>>>>>>", &Result)
	return &Result, nil
}

func (a *Alarm) OpenCurtain(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error) {
	const op = "alarm.OpenDoor"
	log := a.logger.With(
		op, "OpenDoor")
	open := true
	if boolStr == "false" {
		open = false
	}
	message, err := a.client.SmartAlarm().OpenAndCloseCurtain(ctx, &alarm1.OpenAndCloseCurtainReq{
		UserId:     userID,
		DeviceName: deviceName,
		Open:       open,
	})
	if err != nil {
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	log.Info("Called Alarm")

	return entity.StatusMessage{
		Message: message.Message,
	}, nil
}
