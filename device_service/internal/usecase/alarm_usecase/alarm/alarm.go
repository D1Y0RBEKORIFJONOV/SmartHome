package alarm_tv_service_usecase

import (
	"context"
	alarm_entity "device_service/internal/entity/alarm"
)

type AlarmManagementUseCase interface {
	AddSmartAlarm(ctx context.Context, req *alarm_entity.AddSmartAlarmReq) error
	OpenAndCloseCurtain(ctx context.Context, req *alarm_entity.OpenAndCloseCurtainReq) (alarm_entity.OpenAndCloseCurtainRes, error)
	CreateAlarmClock(ctx context.Context, req *alarm_entity.CreateAlarmClockReq) error
	OpenAndClose(ctx context.Context, req *alarm_entity.OpenAndCloseReq) (alarm_entity.OpenAndCloseRes, error)
	GetRemainingTime(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error)
}

type Alarm struct {
	alarm AlarmManagementUseCase
}

func NewAlarmManagementUseCase(alarm AlarmManagementUseCase) Alarm {
	return Alarm{
		alarm: alarm,
	}
}
func (a *Alarm) AddSmartAlarm(ctx context.Context, req *alarm_entity.AddSmartAlarmReq) error {
	return a.alarm.AddSmartAlarm(ctx, req)
}

func (a *Alarm) OpenAndCloseCurtain(ctx context.Context, req *alarm_entity.OpenAndCloseCurtainReq) (alarm_entity.OpenAndCloseCurtainRes, error) {
	return a.alarm.OpenAndCloseCurtain(ctx, req)
}

func (a *Alarm) CreateAlarmClock(ctx context.Context, req *alarm_entity.CreateAlarmClockReq) error {
	return a.alarm.CreateAlarmClock(ctx, req)
}

func (a *Alarm) OpenAndClose(ctx context.Context, req *alarm_entity.OpenAndCloseReq) (alarm_entity.OpenAndCloseRes, error) {
	return a.alarm.OpenAndClose(ctx, req)
}

func (a *Alarm) GetRemainingTime(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error) {
	return a.alarm.GetRemainingTime(ctx, req)
}
