package speaker_entity

type (
	Channel struct {
		ChannelName   string `bson:"channel_name"`
		ChannelNumber string `bson:"channel_number"`
	}

	GetUserChannelRes struct {
		Channels []Channel `bson:"channels"`
		Count    int64     `bson:"count"`
	}

	AddSpeakerReq struct {
		UserID        string `bson:"user_id"`
		ModelName     string `bson:"model_name"`
		CursorChannel uint8  `bson:"cursor_channel"`
		Sound         uint8  `bson:"sound"`
	}

	SpeakerStatusMessage struct {
		Successfully bool `bson:"successfully"`
	}

	AddChannelReq struct {
		UserID      string `bson:"user_id"`
		ChannelName string `bson:"channel_name"`
	}

	GetUserChannelReq struct {
		UserID      string `bson:"user_id"`
		CurrentName string `bson:"current_name"`
	}

	DeleteChannelReq struct {
		UserID      string `bson:"user_id"`
		ChannelName string `bson:"channel_name"`
	}

	DownOrUpVolumeReq struct {
		UserID string `bson:"user_id"`
		Down   bool   `bson:"down"`
		Up     bool   `bson:"up"`
	}

	DownOrUpVolumeRes struct {
		Sound int64 `bson:"sound"`
	}

	PreviousAndNextReq struct {
		UserID string `bson:"user_id"`
		Next   bool   `bson:"next"`
		Back   bool   `bson:"back"`
	}

	PreviousAndNextRes struct {
		Channel *Channel `bson:"channel"`
	}

	Speaker struct {
		UserID        string     `bson:"user_id"`
		ModelName     string     `bson:"model_name"`
		CursorChannel uint8      `bson:"cursor_channel"`
		Sound         uint8      `bson:"sound"`
		Channels      []*Channel `bson:"channels"`
		On            bool       `bson:"on"`
	}

	OnAndOffReq struct {
		UserID string `bson:"user_id"`
		On     bool   `bson:"on"`
		Off    bool   `bson:"off"`
	}
)
