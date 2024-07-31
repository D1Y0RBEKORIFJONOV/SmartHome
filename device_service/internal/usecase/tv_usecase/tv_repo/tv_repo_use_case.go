package tv_repo_use_case

import (
	"context"
	tv_entity "device_service/internal/entity/tv"
)

type (
	Saver interface {
		SaveUsersTV(ctx context.Context, req *tv_entity.TV) error
		SaveChannelsTV(ctx context.Context, req *tv_entity.Channel, userId string) error
	}
	Provider interface {
		GetUsersChannel(ctx context.Context, req *tv_entity.GetUserChannelReq) (*tv_entity.GetUserChannelRes, error)
		GetCursorChannel(ctx context.Context, userID string) (*int, error)
		GetSoundTV(ctx context.Context, userID string) (*int, error)
		IsOnUsersTv(ctx context.Context, userID string) (bool, error)
	}
	Updater interface {
		UpdateSound(ctx context.Context, userID string, val uint8) (*int, error)
		UpdateCursor(ctx context.Context, userID string, val uint8) error
		UpdateTVOn(ctx context.Context, userID string, val bool) error
	}
	Deleter interface {
		DeleteChannel(ctx context.Context, userID, channelName string) error
	}
)

type TVRepoUseCase struct {
	saver    Saver
	provider Provider
	updater  Updater
	deleter  Deleter
}

func NewTVRepoUseCase(saver Saver, provider Provider, updater Updater, deleter Deleter) *TVRepoUseCase {
	return &TVRepoUseCase{
		saver:    saver,
		provider: provider,
		updater:  updater,
		deleter:  deleter,
	}
}

func (t *TVRepoUseCase) SaveUsersTV(ctx context.Context, req *tv_entity.TV) error {
	return t.saver.SaveUsersTV(ctx, req)
}

func (t *TVRepoUseCase) SaveChannelsTV(ctx context.Context, req *tv_entity.Channel, userID string) error {
	return t.saver.SaveChannelsTV(ctx, req, userID)
}

func (t *TVRepoUseCase) GetUsersChannel(ctx context.Context, req *tv_entity.GetUserChannelReq) (*tv_entity.GetUserChannelRes, error) {
	return t.provider.GetUsersChannel(ctx, req)
}

func (t *TVRepoUseCase) GetCursorChannel(ctx context.Context, userID string) (*int, error) {
	return t.provider.GetCursorChannel(ctx, userID)
}

func (t *TVRepoUseCase) UpdateSound(ctx context.Context, userID string, val uint8) (*int, error) {
	return t.updater.UpdateSound(ctx, userID, val)
}

func (t *TVRepoUseCase) UpdateCursor(ctx context.Context, userID string, val uint8) error {
	return t.updater.UpdateCursor(ctx, userID, val)
}

func (t *TVRepoUseCase) DeleteChannel(ctx context.Context, userID, channelName string) error {
	return t.deleter.DeleteChannel(ctx, userID, channelName)
}

func (t *TVRepoUseCase) GetSoundTV(ctx context.Context, userID string) (*int, error) {
	return t.provider.GetSoundTV(ctx, userID)
}

func (t *TVRepoUseCase) IsOnUsersTv(ctx context.Context, userID string) (bool, error) {
	return t.provider.IsOnUsersTv(ctx, userID)
}

func (t *TVRepoUseCase) UpdateTVOn(ctx context.Context, userID string, val bool) error {
	return t.updater.UpdateTVOn(ctx, userID, val)
}
