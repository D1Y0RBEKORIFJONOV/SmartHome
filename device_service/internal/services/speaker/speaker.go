package speaker_service

import (
	"context"
	err_entity "device_service/internal/entity/errors"
	speaker_entity "device_service/internal/entity/speaker"
	"device_service/internal/usecase/speaker_usecase/speaker_repo"
	"errors"
	"fmt"
	"log/slog"
)

type Speaker struct {
	logger  *slog.Logger
	speaker *speaker_repo_use_case.SpeakerRepoUseCase
}

func NewSpeaker(
	logger *slog.Logger,
	speaker *speaker_repo_use_case.SpeakerRepoUseCase) *Speaker {
	return &Speaker{
		logger:  logger,
		speaker: speaker,
	}
}

func (s *Speaker) AddSpeakerToUser(ctx context.Context, req *speaker_entity.AddSpeakerReq) error {
	const op = "Service.Speaker.AddSpeakerToUser"
	log := s.logger.With("method", op)

	log.Info("Called AddSpeakerToUser")
	err := s.speaker.SaveUsersSpeaker(ctx, &speaker_entity.Speaker{
		UserID:        req.UserID,
		ModelName:     req.ModelName,
		CursorChannel: req.CursorChannel,
		Sound:         req.Sound,
		Channels:      GenerateMockPopSongs(),
		On:            false,
	})
	if err != nil {
		log.Error("Failed saving users speaker")
		return err
	}
	log.Info("Saved users speaker")
	return nil
}

func (s *Speaker) AddChannel(ctx context.Context, req *speaker_entity.AddChannelReq) error {
	const op = "Service.Speaker.AddChannel"
	log := s.logger.With("method", op)
	log.Info("Called IsOnUsersSpeaker")
	isOpen, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !isOpen {
		log.Error("User is not on speaker")
		return errors.New("user is not on speaker")
	}
	channels, err := s.speaker.GetUsersChannel(ctx, &speaker_entity.GetUserChannelReq{
		UserID: req.UserID,
	})
	if err != nil {
		log.Error("Failed getting users speaker")
		return err
	}

	log.Info("Called AddChannel")
	err = s.speaker.SaveChannelsSpeaker(ctx, &speaker_entity.Channel{
		ChannelName:   req.ChannelName,
		ChannelNumber: fmt.Sprintf("%02d", channels.Count+1),
	}, req.UserID)
	if err != nil {
		log.Error("Failed saving channels speaker")
		return err
	}
	log.Info("Successfully saved channels speaker")
	return nil
}

func (s *Speaker) GetUserChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error) {
	const op = "Service.Speaker.GetUserChannel"
	log := s.logger.With("method", op)

	log.Info("Called IsOnUsersSpeaker")
	isOpen, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not speaker")
		return nil, errors.New("speaker is closed")
	}

	log.Info("Called GetUserChannel")
	channels, err := s.speaker.GetUsersChannel(ctx, req)
	if err != nil {
		log.Error("Failed getting users speaker")
		return nil, err
	}

	log.Info("Successfully got users speaker channels")
	return channels, nil
}

func (s *Speaker) DeleteChannel(ctx context.Context, req *speaker_entity.DeleteChannelReq) error {
	const op = "Service.Speaker.DeleteChannel"
	log := s.logger.With("method", op)
	log.Info("Called IsOnUsersSpeaker")
	isOpen, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !isOpen {
		log.Error("User is not on speaker")
		return err
	}

	log.Info("Called DeleteChannel")
	err = s.speaker.DeleteChannel(ctx, req.UserID, req.ChannelName)
	if err != nil {
		log.Error("Failed deleting users speaker")
		log.Error(err.Error())
		return err
	}
	log.Info("Successfully deleted users speaker")
	return nil
}

func (s *Speaker) DownOrUpVolume(ctx context.Context, req *speaker_entity.DownOrUpVolumeReq) (*speaker_entity.DownOrUpVolumeRes, error) {
	const op = "Service.Speaker.DownOrUpVoiceSpeaker"
	log := s.logger.With("method", op)
	log.Info("Called IsOnUsersSpeaker")
	isOpen, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not on speaker")
		return nil, errors.New("user is not on speaker")
	}

	log.Info("Called DownOrUpVoiceSpeaker")
	if req.Up == req.Down {
		log.Info("Failed down or up")
		return nil, errors.New("invalid request can't be down and up = false")
	}

	sound, err := s.speaker.GetSoundSpeaker(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users speaker sound")
		return nil, err
	}
	*sound += 1
	if req.Down {
		*sound -= 2
	}

	if req.Up && *sound >= 100 {
		log.Error("Failed getting users speaker sound Err Max limited = 100")
		return &speaker_entity.DownOrUpVolumeRes{}, err_entity.ErrMaxLimitedSound
	}
	if req.Down && *sound <= 0 {
		log.Error("Failed getting users speaker sound Err Min limited = 0")
		return &speaker_entity.DownOrUpVolumeRes{}, err_entity.ErrMinLimitedSound
	}

	sound, err = s.speaker.UpdateSound(ctx, req.UserID, uint8(*sound))
	if err != nil {
		log.Error("Failed updating users speaker sound")
		return nil, err
	}

	return &speaker_entity.DownOrUpVolumeRes{
		Sound: int64(*sound),
	}, nil
}

func (s *Speaker) PreviousAndNext(ctx context.Context, req *speaker_entity.PreviousAndNextReq) (*speaker_entity.PreviousAndNextRes, error) {
	const op = "Service.Speaker.PreviousAndNext"
	log := s.logger.With("method", op)
	log.Info("Called IsOnUsersSpeaker")
	isOpen, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if !isOpen {
		log.Error("User is not on speaker")
		return nil, errors.New("user is not on speaker")
	}
	if req.Next == req.Back {
		log.Info("Failed previous and next req")
		return nil, errors.New("invalid request can't be previous or next false")
	}

	log.Info("Called cursor")
	cursor, err := s.speaker.GetCursorChannel(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users speaker cursor")
		return nil, err
	}
	log.Info("Called all channels")
	channels, err := s.speaker.GetUsersChannel(ctx, &speaker_entity.GetUserChannelReq{
		UserID: req.UserID,
	})
	if err != nil {
		log.Error("Failed getting users speaker channels")
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
	err = s.speaker.UpdateCursor(ctx, req.UserID, uint8(*cursor))
	if err != nil {
		log.Error("Failed updating users speaker cursor")
		return nil, err
	}
	log.Info("Successfully got cursor")

	return &speaker_entity.PreviousAndNextRes{
		Channel: &channels.Channels[*cursor],
	}, nil
}

func (s *Speaker) OnAndOffUsersSpeaker(ctx context.Context, req *speaker_entity.OnAndOffReq) (string, error) {
	const op = "Service.Speaker.OnAndOffUsersSpeaker"
	log := s.logger.With("method", op)
	if req.Off == req.On {
		log.Error("Failed onAndOffUsersSpeaker")
		return "", errors.New("invalid request can't be onAndOffUsersSpeaker = false")
	}
	msg := "opened"
	ok := true
	if req.Off {
		msg = "closed"
		ok = false
	}
	open, err := s.speaker.IsOnUsersSpeaker(ctx, req.UserID)
	if err != nil {
		log.Error("Failed getting users speaker state")
		return "", err
	}

	if open && ok {
		log.Error("Failed speaker already on speaker")
		return "", errors.New("speaker already on speaker")
	}
	if !open && !ok {
		log.Error("Failed speaker already off speaker")
		return "", errors.New("speaker already off speaker")
	}
	err = s.speaker.UpdateSpeakerOn(ctx, req.UserID, ok)
	if err != nil {
		log.Error("Failed updating speaker state")
		return "", err
	}
	return "Successfully: " + msg, nil
}
