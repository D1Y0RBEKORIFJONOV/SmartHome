package alarm_server

import (
	"context"
	alarm_entity "device_service/internal/entity/alarm"
	err_entity "device_service/internal/entity/errors"
	alarm_tv_service_usecase "device_service/internal/usecase/alarm_usecase/alarm"
	"errors"
	alarm1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/smart_alarm"
	"google.golang.org/grpc"
	"time"
)

type Alarm struct {
	alarm1.UnimplementedSmartAlarmServiceServer
	alarm alarm_tv_service_usecase.Alarm
}

func RegisterAlarmServer(grpcServer *grpc.Server, alarm alarm_tv_service_usecase.Alarm) {
	alarm1.RegisterSmartAlarmServiceServer(grpcServer, &Alarm{alarm: alarm})
}

func (a *Alarm) AddSmartAlarm(ctx context.Context, req *alarm1.AddSmartAlarmReq) (*alarm1.AlarmStatusMessage, error) {
	if req.DeviceName == "" || req.ModelName == "" || req.UserId == "" {
		return &alarm1.AlarmStatusMessage{
				Successfully: false,
			}, err_entity.ErrBadRequest{
				Err: errors.New("invalid params"),
			}
	}
	err := a.alarm.AddSmartAlarm(ctx, &alarm_entity.AddSmartAlarmReq{
		DeviceName: req.DeviceName,
		ModelName:  req.ModelName,
		UserID:     req.UserId,
	})
	if err != nil {
		return &alarm1.AlarmStatusMessage{
			Successfully: false,
		}, err
	}
	return &alarm1.AlarmStatusMessage{
		Successfully: true,
	}, nil
}

func (a *Alarm) OpenAndCloseCurtain(ctx context.Context, req *alarm1.OpenAndCloseCurtainReq) (*alarm1.OpenAndCloseCurtainRes, error) {
	if req.UserId == "" || req.DeviceName == "" {
		return nil, err_entity.ErrBadRequest{
			Err: errors.New("invalid params"),
		}
	}
	result, err := a.alarm.OpenAndCloseCurtain(ctx, &alarm_entity.OpenAndCloseCurtainReq{
		UserID:     req.UserId,
		Open:       req.Open,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		return &alarm1.OpenAndCloseCurtainRes{}, err
	}

	return &alarm1.OpenAndCloseCurtainRes{
		Message: result.Message,
	}, nil
}

func (a *Alarm) CreateAlarmClock(ctx context.Context, req *alarm1.CreateAlarmClockReq) (*alarm1.AlarmStatusMessage, error) {
	if req.UserId == "" && req.DeviceName == "" {
		return &alarm1.AlarmStatusMessage{
				Successfully: false,
			}, err_entity.ErrBadRequest{
				Err: errors.New("invalid params"),
			}
	}
	err := a.alarm.CreateAlarmClock(ctx, &alarm_entity.CreateAlarmClockReq{
		UserID:     req.UserId,
		ClockTime:  req.ClockTime,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		return &alarm1.AlarmStatusMessage{
			Successfully: false,
		}, err
	}

	return &alarm1.AlarmStatusMessage{
		Successfully: true,
	}, nil
}

func (a *Alarm) OpenAndClose(ctx context.Context, req *alarm1.OpenAndCloseReq) (*alarm1.OpenAndCloseRes, error) {
	if req.UserId == "" {
		return nil, err_entity.ErrBadRequest{
			Err: errors.New("invalid params"),
		}
	}
	res, err := a.alarm.OpenAndClose(ctx, &alarm_entity.OpenAndCloseReq{
		UserID:     req.UserId,
		Open:       req.Open,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		return &alarm1.OpenAndCloseRes{}, err
	}

	return &alarm1.OpenAndCloseRes{
		Message: res.Message,
	}, nil
}

func (a *Alarm) GetRemainingTime(ctx context.Context, req *alarm1.RemainingTimeReq) (*alarm1.RemainingTimRes, error) {
	if req.UserId == "" || req.DeviceName == "" {
		return nil, err_entity.ErrBadRequest{
			Err: errors.New("invalid params"),
		}
	}
	res, err := a.alarm.GetRemainingTime(ctx, &alarm_entity.RemainingTimeReq{
		UserID:     req.UserId,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		return &alarm1.RemainingTimRes{}, err
	}
	response := alarm1.RemainingTimRes{}
	for _, alarmTemp := range res.Alarms {
		ramaningTime := alarmTemp.AlarmTime.Sub(time.Now())
		if ramaningTime.Seconds() > 0 {
			response.Alarms = append(response.Alarms, &alarm1.Alarm{
				AlarmTime:     alarmTemp.AlarmTime.Format("2006-01-02 15:04:05"),
				RemainingTime: ramaningTime.String(),
			})
		}
	}
	response.Count = res.Count
	return &response, nil
}
