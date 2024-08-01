package speaker_entity

type (
	Channel struct {
		ChannelName   string `bson:"channel_name" json:"channel_name"`
		ChannelNumber string `bson:"channel_number" json:"channel_number"`
	}

	GetUserChannelRes struct {
		Channels []Channel `bson:"channels" json:"channels"`
		Count    int64     `bson:"count" json:"count"`
	}

	AddSpeakerReq struct {
		UserID        string `bson:"user_id" json:"user_id"`
		ModelName     string `bson:"model_name" json:"model_name"`
		CursorChannel uint8  `bson:"cursor_channel" json:"cursor_channel"`
		Sound         uint8  `bson:"sound" json:"sound"`
	}

	SpeakerStatusMessage struct {
		Successfully bool `bson:"successfully" json:"successfully"`
	}

	AddChannelReq struct {
		UserID      string `bson:"user_id" json:"user_id"`
		ChannelName string `bson:"channel_name" json:"channel_name"`
	}

	GetUserChannelReq struct {
		UserID      string `bson:"user_id" json:"user_id"`
		CurrentName string `bson:"current_name" json:"current_name"`
	}

	DeleteChannelReq struct {
		UserID      string `bson:"user_id" json:"user_id"`
		ChannelName string `bson:"channel_name" json:"channel_name"`
	}

	DownOrUpVolumeReq struct {
		UserID string `bson:"user_id" json:"user_id"`
		Down   bool   `bson:"down" json:"down"`
		Up     bool   `bson:"up" json:"up"`
	}

	DownOrUpVolumeRes struct {
		Sound int64 `bson:"sound" json:"sound"`
	}

	PreviousAndNextReq struct {
		UserID string `bson:"user_id" json:"user_id"`
		Next   bool   `bson:"next" json:"next"`
		Back   bool   `bson:"back" json:"back"`
	}

	PreviousAndNextRes struct {
		Channel *Channel `bson:"channel" json:"channel"`
	}

	Speaker struct {
		UserID        string     `bson:"user_id" json:"user_id"`
		ModelName     string     `bson:"model_name" json:"model_name"`
		CursorChannel uint8      `bson:"cursor_channel" json:"cursor_channel"`
		Sound         uint8      `bson:"sound" json:"sound"`
		Channels      []*Channel `bson:"channels" json:"channels"`
		On            bool       `bson:"on" json:"on"`
	}

	OnAndOffReq struct {
		UserID string `bson:"user_id" json:"user_id"`
		On     bool   `bson:"on" json:"on"`
		Off    bool   `bson:"off" json:"off"`
	}
)
