package entity

type (
	Channel struct {
		ChannelName   string `json:"channel_name" redis:"channel_name"`
		ChannelNumber string `json:"channel_number" redis:"channel_number"`
	}

	GetUserChannelRes struct {
		Channels []Channel `json:"channels" redis:"channels"`
		Count    int64     `json:"count" redis:"count"`
	}

	AddTVReq struct {
		UserID    string `json:"-" redis:"user_id"`
		ModelName string `json:"model_name" redis:"model_name"`
	}

	TvStatusMessage struct {
		Successfully bool `json:"successfully" redis:"successfully"`
	}

	AddChannelReq struct {
		UserID      string `json:"-" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
	}

	GetUserChannelReq struct {
		UserID      string `json:"user_id" redis:"user_id"`
		CurrentName string `json:"current_name" redis:"current_name"`
	}

	DeleteChannelReq struct {
		UserID      string `json:"-" redis:"user_id"`
		ChannelName string `json:"channel_name" redis:"channel_name"`
	}

	DownOrUpVoiceTvReq struct {
		UserID string `json:"user_id" redis:"user_id"`
		Down   bool   `json:"down" redis:"down"`
		Up     bool   `json:"up" redis:"up"`
	}

	DownOrUpVoiceTvRes struct {
		Sound int64 `json:"sound" redis:"sound"`
	}

	PreviousAndNextReq struct {
		UserID string `json:"user_id" redis:"user_id"`
		Next   bool   `json:"next" redis:"next"`
		Back   bool   `json:"back" redis:"back"`
	}

	PreviousAndNextRes struct {
		Channel *Channel `json:"channel" redis:"channel"`
	}

	TV struct {
		UserID        string     `json:"user_id" redis:"user_id"`
		ModelName     string     `json:"model_name" redis:"model_name"`
		CursorChannel uint8      `json:"cursor_channel" redis:"cursor_channel"`
		Sound         uint8      `json:"sound" redis:"sound"`
		Channels      []*Channel `json:"channels" redis:"channels"`
		On            bool       `json:"on" redis:"on"`
	}

	OnAndOfReq struct {
		UserID string `json:"user_id" redis:"user_id"`
		On     bool   `json:"on" redis:"on"`
		Off    bool   `json:"off" redis:"off"`
	}
)
