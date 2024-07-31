package alarm_repo

import (
	"context"
	alarm_entity "device_service/internal/entity/alarm"
)

type (
	Saver interface {
		SaveAlarm(ctx context.Context, alarm alarm_entity.SmartAlarm) error
		SaveAlarmClock(ctx context.Context, req *alarm_entity.Alarm, userId, deviceName string) error
	}
	Updater interface {
		UpdateCurtain(ctx context.Context, userId, deviceName string, val bool) error
		UpdateDoor(ctx context.Context, userId, deviceName string, val bool) error
	}
	Provider interface {
		GetAlarmUser(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error)
		IsDeviceExists(ctx context.Context, userId, deviceName string) (bool, error)
	}
)

type AlarmRepoUseCase struct {
	saver    Saver
	updater  Updater
	provider Provider
}

func NewAlarmUseCase(saver Saver, updater Updater, provider Provider) *AlarmRepoUseCase {
	return &AlarmRepoUseCase{
		saver:    saver,
		updater:  updater,
		provider: provider,
	}
}

func (a *AlarmRepoUseCase) SaveAlarm(ctx context.Context, alarm alarm_entity.SmartAlarm) error {
	return a.saver.SaveAlarm(ctx, alarm)
}

func (a *AlarmRepoUseCase) SaveAlarmClock(ctx context.Context, req *alarm_entity.Alarm, userId, deviceName string) error {
	return a.saver.SaveAlarmClock(ctx, req, userId, deviceName)
}

func (a *AlarmRepoUseCase) IsDeviceExists(ctx context.Context, userId, deviceName string) (bool, error) {
	return a.provider.IsDeviceExists(ctx, userId, deviceName)
}
func (a *AlarmRepoUseCase) UpdateCurtain(ctx context.Context, userId, deviceName string, val bool) error {
	return a.updater.UpdateCurtain(ctx, userId, deviceName, val)
}

func (a *AlarmRepoUseCase) UpdateDoor(ctx context.Context, userId, deviceName string, val bool) error {
	return a.updater.UpdateDoor(ctx, userId, deviceName, val)
}

func (a *AlarmRepoUseCase) GetAlarmUser(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error) {
	return a.provider.GetAlarmUser(ctx, req)
}
