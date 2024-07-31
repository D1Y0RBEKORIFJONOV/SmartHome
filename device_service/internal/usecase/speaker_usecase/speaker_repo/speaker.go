package speaker_repo_use_case

import (
	"context"
	speaker_entity "device_service/internal/entity/speaker"
)

type (
	Saver interface {
		SaveUsersSpeaker(ctx context.Context, req *speaker_entity.Speaker) error
		SaveChannelsSpeaker(ctx context.Context, req *speaker_entity.Channel, userID string) error
	}
	Provider interface {
		GetUsersChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error)
		GetCursorChannel(ctx context.Context, userID string) (*int, error)
		GetSoundSpeaker(ctx context.Context, userID string) (*int, error)
		IsOnUsersSpeaker(ctx context.Context, userID string) (bool, error)
	}
	Updater interface {
		UpdateSound(ctx context.Context, userID string, val uint8) (*int, error)
		UpdateCursor(ctx context.Context, userID string, val uint8) error
		UpdateSpeakerOn(ctx context.Context, userID string, val bool) error
	}
	Deleter interface {
		DeleteChannel(ctx context.Context, userID, channelName string) error
	}
)

type SpeakerRepoUseCase struct {
	saver    Saver
	provider Provider
	updater  Updater
	deleter  Deleter
}

func NewSpeakerRepoUseCase(saver Saver, provider Provider, updater Updater, deleter Deleter) *SpeakerRepoUseCase {
	return &SpeakerRepoUseCase{
		saver:    saver,
		provider: provider,
		updater:  updater,
		deleter:  deleter,
	}
}

func (t *SpeakerRepoUseCase) SaveUsersSpeaker(ctx context.Context, req *speaker_entity.Speaker) error {
	return t.saver.SaveUsersSpeaker(ctx, req)
}

func (t *SpeakerRepoUseCase) SaveChannelsSpeaker(ctx context.Context, req *speaker_entity.Channel, userID string) error {
	return t.saver.SaveChannelsSpeaker(ctx, req, userID)
}

func (t *SpeakerRepoUseCase) GetUsersChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error) {
	return t.provider.GetUsersChannel(ctx, req)
}

func (t *SpeakerRepoUseCase) GetCursorChannel(ctx context.Context, userID string) (*int, error) {
	return t.provider.GetCursorChannel(ctx, userID)
}

func (t *SpeakerRepoUseCase) UpdateSound(ctx context.Context, userID string, val uint8) (*int, error) {
	return t.updater.UpdateSound(ctx, userID, val)
}

func (t *SpeakerRepoUseCase) UpdateCursor(ctx context.Context, userID string, val uint8) error {
	return t.updater.UpdateCursor(ctx, userID, val)
}

func (t *SpeakerRepoUseCase) DeleteChannel(ctx context.Context, userID, channelName string) error {
	return t.deleter.DeleteChannel(ctx, userID, channelName)
}

func (t *SpeakerRepoUseCase) GetSoundSpeaker(ctx context.Context, userID string) (*int, error) {
	return t.provider.GetSoundSpeaker(ctx, userID)
}

func (t *SpeakerRepoUseCase) IsOnUsersSpeaker(ctx context.Context, userID string) (bool, error) {
	return t.provider.IsOnUsersSpeaker(ctx, userID)
}

func (t *SpeakerRepoUseCase) UpdateSpeakerOn(ctx context.Context, userID string, val bool) error {
	return t.updater.UpdateSpeakerOn(ctx, userID, val)
}
