package alarm_usecase

import (
	"api_gate_way/internal/entity"
	"context"
)

type AlarmUseCase interface {
	AddSmartAlarm(ctx context.Context, alarm entity.AddSmartAlarmReq) (entity.StatusMessage, error)
	OpenCurtain(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error)
	CreateAlarmClock(ctx context.Context, req *entity.CreateAlarmClockReq) (entity.StatusMessage, error)
	OpenDoor(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error)
	GetRemainingTime(ctx context.Context, userID, deviceName string) (*entity.RemainingTimRes, error)
}
type Alarm struct {
	alarm AlarmUseCase
}

func NewAlarm(alarm AlarmUseCase) *Alarm {
	return &Alarm{
		alarm: alarm,
	}
}

func (a *Alarm) AddSmartAlarm(ctx context.Context, alarm entity.AddSmartAlarmReq) (entity.StatusMessage, error) {
	return a.alarm.AddSmartAlarm(ctx, alarm)
}
func (a *Alarm) OpenCurtain(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error) {
	return a.alarm.OpenCurtain(ctx, userID, boolStr, deviceName)

}
func (a *Alarm) CreateAlarmClock(ctx context.Context, req *entity.CreateAlarmClockReq) (entity.StatusMessage, error) {
	return a.alarm.CreateAlarmClock(ctx, req)
}
func (a *Alarm) OpenDoor(ctx context.Context, userID, boolStr, deviceName string) (entity.StatusMessage, error) {
	return a.alarm.OpenDoor(ctx, userID, boolStr, deviceName)
}

func (a *Alarm) GetRemainingTime(ctx context.Context, userID, deviceName string) (*entity.RemainingTimRes, error) {
	return a.alarm.GetRemainingTime(ctx, userID, deviceName)
}
