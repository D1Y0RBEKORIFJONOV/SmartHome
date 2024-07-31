package alarm_entity

import "time"

type (
	AlarmStatusMessage struct {
		Successfully bool `bson:"successfully"`
	}

	AddSmartAlarmReq struct {
		UserID     string `bson:"user_id"`
		DeviceName string `bson:"device_name"`
		ModelName  string `bson:"model_name"`
	}

	OpenAndCloseCurtainReq struct {
		UserID     string `bson:"user_id"`
		Open       bool   `bson:"open"`
		DeviceName string `bson:"device_name"`
	}

	OpenAndCloseCurtainRes struct {
		Message string `bson:"message"`
	}

	CreateAlarmClockReq struct {
		UserID     string `bson:"user_id"`
		ClockTime  string `bson:"clock_time"`
		DeviceName string `bson:"device_name"`
	}

	OpenAndCloseReq struct {
		UserID     string `bson:"user_id"`
		Open       bool   `bson:"open"`
		DeviceName string `bson:"device_name"`
	}

	OpenAndCloseRes struct {
		Message string `bson:"message"`
	}

	RemainingTimeReq struct {
		UserID     string `bson:"user_id"`
		DeviceName string `bson:"device_name"`
	}

	Alarm struct {
		AlarmTime     time.Time `bson:"alarm_time"`
		RemainingTime time.Time `bson:"remaining_time"`
	}

	RemainingTimRes struct {
		Alarms []Alarm `bson:"alarms"`
		Count  int64   `bson:"count"`
	}
	SmartAlarm struct {
		UserID      string  `bson:"user_id"`
		DeviceName  string  `bson:"device_name"`
		ModelName   string  `bson:"model_name"`
		Alarms      []Alarm `bson:"alarms"`
		OpenDoor    bool    `bson:"door"`
		CurtainOpen bool    `bson:"curtain"`
	}
)
