package alarm_service

import (
	"context"
	alarm_entity "device_service/internal/entity/alarm"
	err_entity "device_service/internal/entity/errors"
	"device_service/internal/usecase/alarm_usecase/alarm_repo"
	"fmt"
	"log/slog"
	"time"
)

type Alarm struct {
	logger *slog.Logger
	alarm  alarm_repo.AlarmRepoUseCase
}

func NewAlarm(alarm alarm_repo.AlarmRepoUseCase, logger *slog.Logger) *Alarm {
	return &Alarm{
		alarm:  alarm,
		logger: logger,
	}
}

func (a *Alarm) AddSmartAlarm(ctx context.Context, req *alarm_entity.AddSmartAlarmReq) error {
	const op = "Service.Alarm.AddSmartAlarm"
	log := a.logger.With(
		"method", op)
	log.Info("Called IsDeviceExists")
	IsExists, err := a.alarm.IsDeviceExists(ctx, req.UserID, req.DeviceName)
	if err != nil {
		log.Error(err.Error())
		return err_entity.ErrorAlreadyExists
	}
	if IsExists {
		log.Error("Err: device already exists")
		return err_entity.ErrorNotFound
	}
	log.Info("Called SaveSmartAlarm")
	err = a.alarm.SaveAlarm(ctx, alarm_entity.SmartAlarm{
		UserID:      req.UserID,
		DeviceName:  req.DeviceName,
		ModelName:   req.ModelName,
		Alarms:      []alarm_entity.Alarm{},
		OpenDoor:    false,
		CurtainOpen: false,
	})
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (a *Alarm) OpenAndCloseCurtain(ctx context.Context, req *alarm_entity.OpenAndCloseCurtainReq) (alarm_entity.OpenAndCloseCurtainRes, error) {
	const op = "Service.Alarm.OpenAndCloseCurtain"
	log := a.logger.With(
		"method", op)

	log.Info("Called UpdateCurtain")
	err := a.alarm.UpdateCurtain(ctx, req.UserID, req.DeviceName, req.Open)
	if err != nil {
		log.Error(err.Error())
		return alarm_entity.OpenAndCloseCurtainRes{}, err
	}
	message := "curtain is successfully opened"
	if !req.Open {
		message = "curtain is successfully closed"
	}
	return alarm_entity.OpenAndCloseCurtainRes{
		Message: message,
	}, nil
}

func addTimeToToday(timeStr string) (*time.Time, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return &time.Time{}, err
	}

	now := time.Now()

	newTime := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
	fmt.Println(newTime, ">>>>>>>>>>>>>>>>")
	return &newTime, nil
}

func (a *Alarm) CreateAlarmClock(ctx context.Context, req *alarm_entity.CreateAlarmClockReq) error {
	const op = "Service.Alarm.CreateAlarmClock"
	log := a.logger.With(
		"method", op)

	timeNow, err := addTimeToToday(req.ClockTime)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("Called SaveAlarmClock")
	err = a.alarm.SaveAlarmClock(ctx, &alarm_entity.Alarm{
		AlarmTime:     *timeNow,
		RemainingTime: *timeNow,
	}, req.UserID, req.DeviceName)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (a *Alarm) OpenAndClose(ctx context.Context, req *alarm_entity.OpenAndCloseReq) (alarm_entity.OpenAndCloseRes, error) {
	const op = "Service.Alarm.OpenAndClose"
	log := a.logger.With(
		"method", op)

	log.Info("Called UpdateDoor")
	err := a.alarm.UpdateDoor(ctx, req.UserID, req.DeviceName, req.Open)
	if err != nil {
		log.Error(err.Error())
		return alarm_entity.OpenAndCloseRes{}, err
	}
	message := "door is successfully opened"
	if !req.Open {
		message = "door is successfully closed"
	}
	return alarm_entity.OpenAndCloseRes{
		Message: message,
	}, nil
}

func (a *Alarm) GetRemainingTime(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error) {
	const op = "Service.Alarm.GetRemainingTime"
	log := a.logger.With(
		"method", op)
	log.Info("Called GetRemainingTime")
	res, err := a.alarm.GetAlarmUser(ctx, req)
	if err != nil {
		log.Error(err.Error())
		return alarm_entity.RemainingTimRes{}, err
	}
	return res, nil
}
