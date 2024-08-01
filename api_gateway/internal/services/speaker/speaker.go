package speaker_service

import (
	"api_gate_way/internal/config"
	"api_gate_way/internal/entity"
	clientgrpcserver "api_gate_way/internal/infastructure/client_grpc_server"
	"api_gate_way/internal/infastructure/producer"
	"context"
	"encoding/json"
	"errors"
	speaker1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/speaker"
	"log"
	"log/slog"
)

type Speaker struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
	p      *producer.Producer
}

func New(logger *slog.Logger, client clientgrpcserver.ServiceClient) *Speaker {
	cfg := config.New()
	p, err := producer.NewProducer(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &Speaker{
		logger: logger,
		client: client,
		p:      p,
	}
}

func (s *Speaker) AddSpeaker(ctx context.Context, req *entity.AddSpeakerReq) (*entity.StatusMessage, error) {
	const op = "speaker.AddSpeaker"
	log := s.logger.With("method", op)
	reqSpeaker := entity.AddTVReqToPub{
		UserID:     req.UserID,
		MethodName: "SPEAKER.AddSpeaker",
		ModelName:  req.ModelName,
	}
	reqBody, err := json.Marshal(reqSpeaker)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	err = s.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Info("Added Speaker successfully")
	return &entity.StatusMessage{
		Message: "Added Speaker in processing",
	}, nil
}

func (s *Speaker) OpenSpeaker(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error) {
	const op = "speaker.OpenSpeaker"
	log := s.logger.With("method", op)
	log.Info("Opening Speaker")
	if boolStr != "true" && boolStr != "false" {
		return nil, errors.New("bad request")
	}
	open := true
	off := false
	if boolStr == "false" {
		open = false
		off = true
	}
	message, err := s.client.SpeakerService().OnAndOffUserSpeaker(ctx, &speaker1.OnAndOffUserSpeakerReq{
		UserId: id,
		On:     open,
		Off:    off,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Open Speaker successfully")

	return &entity.StatusMessage{
		Message: message.Message,
	}, nil
}

func (s *Speaker) GetChannels(ctx context.Context, id string) (*entity.GetUserSongsRes, error) {
	const op = "speaker.GetChannels"
	log := s.logger.With("method", op)
	log.Info("Getting Channels")
	channels, err := s.client.SpeakerService().GetUserChannel(ctx, &speaker1.GetUserChannelReqS{
		UserId: id,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var result entity.GetUserSongsRes
	result.Count = channels.Count
	for _, ch := range channels.Channels {
		result.Channels = append(result.Channels, entity.Channel{
			ChannelName:   ch.ChannelName,
			ChannelNumber: ch.ChannelNumber,
		})
	}
	return &result, nil
}

func (s *Speaker) NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error) {
	const op = "speaker.NextChannel"
	log := s.logger.With("method", op)
	log.Info("Getting Channels")
	next := true
	back := false
	if boolStr == "false" {
		next = false
		back = true
	}

	channel, err := s.client.SpeakerService().PreviousAndNext(ctx, &speaker1.PreviousAndNextReqS{
		UserId: id,
		Next:   next,
		Back:   back,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &entity.PreviousAndNextRes{
		Channel: &entity.Channel{
			ChannelName:   channel.Channel.ChannelName,
			ChannelNumber: channel.Channel.ChannelNumber,
		},
	}, nil
}

func (s *Speaker) AddChannel(ctx context.Context, req *entity.AddChannelReqChannel) (*entity.StatusMessage, error) {
	const op = "speaker.AddChannel"
	log := s.logger.With("method", op)
	log.Info("Adding Channel")

	addReq := entity.AddChannelReqPub{
		UserID:      req.UserID,
		ChannelName: req.ChannelName,
		MethodName:  "Speaker.AddChannel",
	}
	reqBody, err := json.Marshal(addReq)
	if err != nil {
		log.Error(err.Error())
	}
	err = s.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Added Channel successfully")
	return &entity.StatusMessage{
		Message: "Added Channel in processing",
	}, nil
}

func (s *Speaker) DeleteChannel(ctx context.Context, req *entity.DeleteChannelReqChannel) (*entity.StatusMessage, error) {
	const op = "speaker.DeleteChannel"
	log := s.logger.With("method", op)
	log.Info("Deleting Channel")
	deleteReq := entity.DeleteChannelReqPub{
		UserID:      req.UserID,
		ChannelName: req.ChannelName,
		MethodName:  "Speaker.DeleteChannel",
	}
	reqBody, err := json.Marshal(deleteReq)
	if err != nil {
		log.Error(err.Error())
	}
	err = s.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Deleted Channel successfully")
	return &entity.StatusMessage{
		Message: "Deleted Channel in processing",
	}, nil
}

func (s *Speaker) UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceSpeakerRes, error) {
	const op = "speaker.UpVoice"
	log := s.logger.With("method", op)
	up := true
	down := false
	if boolStr == "false" {
		up = false
		down = true
	}
	res, err := s.client.SpeakerService().DownOrUpVolume(ctx, &speaker1.DownOrUpVolumeReqS{
		UserId: userID,
		Down:   down,
		Up:     up,
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Up Voice successfully")
	return &entity.DownOrUpVoiceSpeakerRes{
		Sound: res.Sound,
	}, nil
}
