package tv_service

import (
	"context"
	err_entity "device_service/internal/entity/errors"
	tv_entity "device_service/internal/entity/tv"
	"device_service/internal/usecase/tv_usecase/tv_repo"
	"errors"
	"fmt"
	"log/slog"
)

type TV struct {
	logger *slog.Logger
	tv     *tv_repo_use_case.TVRepoUseCase
}

func NewTV(
	logger *slog.Logger,
	tv *tv_repo_use_case.TVRepoUseCase) *TV {
	return &TV{
		logger: logger,
		tv:     tv,
	}
}

func (t *TV) AddTvToUser(ctx context.Context, req *tv_entity.AddTVReq) error {
	const op = "Service.TV.AddTvToUser"
	log := t.logger.With("method", op)

	log.Info("Called AddTvToUser")
	err := t.tv.SaveUsersTV(ctx, &tv_entity.TV{
		UserID:        req.UserID,
		ModelName:     req.ModelName,
		CursorChannel: req.CursorChannel,
		Sound:         req.Sound,
		Channels:      GenerateMockChannels(),
		On:            false,
	})
	if err != nil {
		log.Error("Failed saving users TV")
		return err
	}
	log.Info("Saved users TV")
	return nil
}

func (t *TV) AddChannel(ctx context.Context, req *tv_entity.AddChannelReq) error {
	const op = "Service.TV.AddChannel"
	log := t.logger.With("method", op)
	log.Info("Called IsOnUsersTV")
	isOpen, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !isOpen {
		log.Error("User is not on TV")
		return errors.New("user is not on TV")
	}
	channels, err := t.tv.GetUsersChannel(ctx, &tv_entity.GetUserChannelReq{
		UserID: req.UserID,
	})
	if err != nil {
		log.Error("Failed getting users TV")
		return err
	}

	log.Info("Called AddChannel")
	err = t.tv.SaveChannelsTV(ctx, &tv_entity.Channel{
		ChannelName:   req.ChannelName,
		ChannelNumber: fmt.Sprintf("%02d", channels.Count+1),
	}, req.UserID)
	if err != nil {
		log.Error("Failed saving channels TV")
		return err
	}
	log.Info("Successfully saved channels TV")
	return nil
}

func (t *TV) GetUserChannel(ctx context.Context, req *tv_entity.GetUserChannelReq) (*tv_entity.GetUserChannelRes, error) {
	const op = "Service.TV.GetUserChannel"
	log := t.logger.With("method", op)

	log.Info("Called IsOnUsersTV")
	isOpen, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not  TV")
		return nil, errors.New("TV is closed")
	}

	log.Info("Called GetUserChannel")
	channels, err := t.tv.GetUsersChannel(ctx, req)
	if err != nil {
		log.Error("Failed getting users TV")
		return nil, err
	}

	log.Info("Successfully got users TV channels")
	return channels, nil
}
func (t *TV) DeleteChannel(ctx context.Context, req *tv_entity.DeleteChannelReq) error {
	const op = "Service.TV.DeleteChannel"
	log := t.logger.With("method", op)
	log.Info("Called IsOnUsersTV")
	isOpen, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !isOpen {
		log.Error("User is not on TV")
		return err
	}

	log.Info("Called DeleteChannel")
	err = t.tv.DeleteChannel(ctx, req.UserID, req.ChannelName)
	if err != nil {
		log.Error("Failed deleting users TV")
		log.Error(err.Error())
		return err
	}
	log.Info("Successfully deleted users TV")
	return nil
}

func (t *TV) DownOrUpVVoiceTv(ctx context.Context, req *tv_entity.DownOrUpVoiceTvReq) (*tv_entity.DownOrUpVoiceTvRes, error) {
	const op = "Service.TV.DownOrUpVoiceTv"
	log := t.logger.With("method", op)
	log.Info("Called IsOnUsersTV")
	isOpen, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not on TV")
		return nil, errors.New("user is not on TV")
	}

	log.Info("Called DownOrUpVoiceTv")
	if req.Up == req.Down {
		log.Info("Failed down or up")
		return nil, errors.New("invalid request can't been down and up = false")
	}

	sound, err := t.tv.GetSoundTV(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users TV sound")
		return nil, err
	}
	*sound += 1
	if req.Down {
		*sound -= 2
	}

	if req.Up && *sound >= 100 {
		log.Error("Failed getting users TV sound Err Max limited = 100")
		return &tv_entity.DownOrUpVoiceTvRes{}, err_entity.ErrMaxLimitedSound
	}
	if req.Down && *sound <= 0 {
		log.Error("Failed getting users TV sound Err Min limited = 0")
		return &tv_entity.DownOrUpVoiceTvRes{}, err_entity.ErrMinLimitedSound
	}

	sound, err = t.tv.UpdateSound(ctx, req.UserID, uint8(*sound))
	if err != nil {
		log.Error("Failed getting users TV sound")
		return nil, err
	}

	return &tv_entity.DownOrUpVoiceTvRes{
		Sound: int64(*sound),
	}, nil
}

func (t *TV) PreviousAndNext(ctx context.Context, req *tv_entity.PreviousAndNextReq) (*tv_entity.PreviousAndNextRes, error) {
	const op = "Service.TV.PreviousAndNext"
	log := t.logger.With("method", op)
	log.Info("Called IsOnUsersTV")
	isOpen, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not on TV")
		return nil, errors.New("user is not on TV")
	}
	if req.Next == req.Back {
		log.Info("Failed previous and next req")
		return nil, errors.New("invalid request can't be previous or next false")
	}

	log.Info("Called cursor")
	cursor, err := t.tv.GetCursorChannel(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users TV cursor")
		return nil, err
	}
	log.Info("Called all channels")
	channels, err := t.tv.GetUsersChannel(ctx, &tv_entity.GetUserChannelReq{
		UserID: req.UserID,
	})
	if err != nil {
		log.Error("Failed getting users TV channels")
		return nil, err
	}
	*cursor += 1
	if req.Back {
		*cursor -= 2
	}

	if req.Next && *cursor > len(channels.Channels)-1 {
		*cursor = 0
	}
	if req.Back && int(*cursor) < 0 {
		*cursor = len(channels.Channels) - 1
	}

	log.Info("Called cursor")
	err = t.tv.UpdateCursor(ctx, req.UserID, uint8(*cursor))
	if err != nil {
		log.Error("Failed getting users TV cursor")
		return nil, err
	}
	log.Info("Successfully got cursor")

	return &tv_entity.PreviousAndNextRes{
		Channel: &channels.Channels[*cursor],
	}, nil
}

func (t *TV) OnAndOffUsersTv(ctx context.Context, req *tv_entity.OnAndOfReq) (string, error) {
	const op = "Service.TV.OnAndOffUsersTv"
	log := t.logger.With("method", op)
	if req.Off == req.On {
		log.Error("Failed onAndOffUsersTv")
		return "", errors.New("invalid request can't be onAndOffUsersTv = false")
	}
	msg := "opened"
	ok := true
	if req.Off {
		msg = "closed"
		ok = false
	}
	open, err := t.tv.IsOnUsersTv(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users TV sound")
		return "", err
	}

	if open && ok {
		log.Error("Failed tv already on TV")
		return "", errors.New("tv already on TV")
	}
	if !open && !ok {
		log.Error("Failed tv already off TV")
		return "", errors.New("tv already off TV")
	}
	err = t.tv.UpdateTVOn(ctx, req.UserID, ok)
	if err != nil {
		log.Error("Failed updating TV sound")
		return "", err
	}
	return "Successfully: " + msg, nil
}
