package entity

type (
	CreateUserReqToPub struct {
		FirstName  string `json:"first_name" redis:"first_name"`
		LastName   string `json:"last_name" redis:"last_name"`
		Email      string `json:"email" redis:"email"`
		Address    string `json:"address" redis:"address"`
		Password   string `json:"password" redis:"password"`
		MethodName string `json:"method_name" redis:"method_name"`
	}
	UpdateUserReqToPub struct {
		UserID     string `json:"id" redis:"id"`
		FirstName  string `json:"first_name" redis:"first_name"`
		LastName   string `json:"last_name" redis:"last_name"`
		MethodName string `json:"method_name" redis:"method_name"`
	}
	UpdateEmailReqToPub struct {
		UserID     string `json:"id" redis:"id"`
		NewEmail   string `json:"new_email" redis:"new_email"`
		MethodName string `json:"method_name" redis:"method_name"`
	}
	UpdatePasswordReqToPub struct {
		UserID      string `json:"id" redis:"id"`
		Password    string `json:"password" redis:"password"`
		NewPassword string `json:"new_password" redis:"new_password"`
		MethodName  string `json:"method_name" redis:"method_name"`
	}
	DeleteUserReqToPub struct {
		UserID        string `json:"id" redis:"id"`
		IsHardDeleted bool   `json:"is_hard_deleted" redis:"is_hard_deleted"`
		MethodName    string `json:"method_name" redis:"method_name"`
	}
	AddTVReqToPub struct {
		UserID     string `json:"user_id" redis:"user_id"`
		ModelName  string `json:"model_name" redis:"model_name"`
		MethodName string `json:"method_name" redis:"method_name"`
	}
	AddChannelReqPub struct {
		UserID      string `json:"user_id" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
		MethodName  string `json:"method_name" redis:"method_name"`
	}
	DeleteChannelReqPub struct {
		UserID      string `json:"user_id" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
		MethodName  string `json:"method_name" redis:"method_name"`
	}

	AddSmartAlarmReqPub struct {
		UserID     string `bson:"user_id" json:"user_id"`
		DeviceName string `bson:"device_name" json:"device_name"`
		ModelName  string `bson:"model_name" json:"model_name"`
		MethodName string `bson:"method_name" json:"method_name"`
	}

	CreateAlarmClockReqPub struct {
		UserID     string `bson:"user_id" json:"user_id"`
		ClockTime  string `bson:"clock_time" json:"clock_time"`
		DeviceName string `bson:"device_name" json:"device_name"`
		MethodName string `bson:"method_name" json:"method_name"`
	}
)
