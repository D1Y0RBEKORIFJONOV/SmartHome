package entity

type (
	Songs struct {
		ChannelName   string `json:"channel_name" redis:"channel_name"`
		ChannelNumber string `json:"channel_number" redis:"channel_number"`
	}

	GetUserSongsRes struct {
		Channels []Channel `json:"channels" redis:"channels"`
		Count    int64     `json:"count" redis:"count"`
	}

	AddSpeakerReq struct {
		UserID    string `json:"-" redis:"user_id"`
		ModelName string `json:"model_name" redis:"model_name"`
	}

	SpeakerStatusMessage struct {
		Successfully bool `json:"successfully" redis:"successfully"`
	}

	AddChannelReqChannel struct {
		UserID      string `json:"-" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
	}

	GetUserChannelReqChannel struct {
		UserID      string `json:"user_id" redis:"user_id"`
		CurrentName string `json:"current_name" redis:"current_name"`
	}

	DeleteChannelReqChannel struct {
		UserID      string `json:"-" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
	}

	DownOrUpVoiceSpeakerReq struct {
		UserID string `json:"user_id" redis:"user_id"`
		Down   bool   `json:"down" redis:"down"`
		Up     bool   `json:"up" redis:"up"`
	}

	DownOrUpVoiceSpeakerRes struct {
		Sound int64 `json:"sound" redis:"sound"`
	}

	Speaker struct {
		UserID        string     `json:"user_id" redis:"user_id"`
		ModelName     string     `json:"model_name" redis:"model_name"`
		CursorChannel uint8      `json:"cursor_channel" redis:"cursor_channel"`
		Sound         uint8      `json:"sound" redis:"sound"`
		Channels      []*Channel `json:"channels" redis:"channels"`
		On            bool       `json:"on" redis:"on"`
	}

	OnAndOffReq struct {
		UserID string `json:"user_id" redis:"user_id"`
		On     bool   `json:"on" redis:"on"`
		Off    bool   `json:"off" redis:"off"`
	}
)
