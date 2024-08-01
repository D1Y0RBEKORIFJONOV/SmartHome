package entity

type (
	AlarmStatusMessage struct {
		Successfully bool `bson:"successfully" json:"successfully"`
	}

	AddSmartAlarmReq struct {
		UserID     string `bson:"-" json:"-"`
		DeviceName string `bson:"device_name" json:"device_name"`
		ModelName  string `bson:"model_name" json:"model_name"`
	}

	OpenAndCloseCurtainReq struct {
		UserID     string `bson:"user_id" json:"user_id"`
		Open       bool   `bson:"open" json:"open"`
		DeviceName string `bson:"device_name" json:"device_name"`
	}

	OpenAndCloseCurtainRes struct {
		Message string `bson:"message" json:"message"`
	}

	CreateAlarmClockReq struct {
		UserID     string `bson:"-" json:"-"`
		ClockTime  string `bson:"clock_time" json:"clock_time"`
		DeviceName string `bson:"device_name" json:"device_name"`
	}

	OpenAndCloseReq struct {
		UserID     string `bson:"user_id" json:"user_id"`
		Open       bool   `bson:"open" json:"open"`
		DeviceName string `bson:"device_name" json:"device_name"`
	}

	OpenAndCloseRes struct {
		Message string `bson:"message" json:"message"`
	}

	RemainingTimeReq struct {
		UserID     string `bson:"user_id" json:"user_id"`
		DeviceName string `bson:"device_name" json:"device_name"`
	}

	Alarm struct {
		AlarmTime     string `bson:"alarm_time" json:"alarm_time"`
		RemainingTime string `bson:"remaining_time" json:"remaining_time"`
	}

	RemainingTimRes struct {
		Alarms []Alarm `bson:"alarms" json:"alarms"`
		Count  int64   `bson:"count" json:"count"`
	}

	SmartAlarm struct {
		UserID      string  `bson:"user_id" json:"user_id"`
		DeviceName  string  `bson:"device_name" json:"device_name"`
		ModelName   string  `bson:"model_name" json:"model_name"`
		Alarms      []Alarm `bson:"alarms" json:"alarms"`
		OpenDoor    bool    `bson:"door" json:"door"`
		CurtainOpen bool    `bson:"curtain" json:"curtain"`
	}
)
