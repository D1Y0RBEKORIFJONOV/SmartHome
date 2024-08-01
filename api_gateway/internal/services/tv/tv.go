package tv_service

import (
	"api_gate_way/internal/config"
	"api_gate_way/internal/entity"
	clientgrpcserver "api_gate_way/internal/infastructure/client_grpc_server"
	"api_gate_way/internal/infastructure/producer"
	"context"
	"encoding/json"
	"errors"
	tv1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/TV"
	"log"
	"log/slog"
)

type TV struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
	p      *producer.Producer
}

func New(logger *slog.Logger, client clientgrpcserver.ServiceClient) *TV {
	cfg := config.New()
	p, err := producer.NewProducer(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &TV{
		logger: logger,
		client: client,
		p:      p,
	}
}

func (t *TV) AddTV(ctx context.Context, req *entity.AddTVReq) (*entity.StatusMessage, error) {
	const op = "tv.AddTV"
	log := t.logger.With("method", op)
	reqTv := entity.AddTVReqToPub{
		UserID:     req.UserID,
		MethodName: "TV.AddTvToUser",
		ModelName:  req.ModelName,
	}
	reqBody, err := json.Marshal(reqTv)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	err = t.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Info("Added TV successfully")
	return &entity.StatusMessage{
		Message: "Added TV in processing",
	}, nil
}

func (t *TV) OpenTV(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error) {
	const op = "tv.OpenTV"
	log := t.logger.With("method", op)
	log.Info("Opening TV")
	if boolStr != "true" && boolStr != "false" {
		return nil, errors.New(
			"bad request")
	}
	open1 := true
	close1 := false
	if boolStr == "false" {
		open1 = false
		close1 = true
	}
	message, err := t.client.TvService().OnAndOffUserTv(ctx, &tv1.OnAndOffUserTvReq{
		UserId: id,
		On:     open1,
		Off:    close1,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Open TV successfully")

	return &entity.StatusMessage{
		Message: message.Message,
	}, nil
}

func (t *TV) GetChannels(ctx context.Context, id string) (*entity.GetUserChannelRes, error) {
	const op = "tv.GetChannels"
	log := t.logger.With("method", op)
	log.Info("Getting Channels")
	channesl, err := t.client.TvService().GetUserChannel(ctx, &tv1.GetUserChannelReq{
		UserId: id,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var result entity.GetUserChannelRes
	result.Count = channesl.Count
	for _, ch := range channesl.Channels {
		result.Channels = append(result.Channels, entity.Channel{
			ChannelName:   ch.ChannelName,
			ChannelNumber: ch.ChannelNumber,
		})
	}
	return &result, nil
}

func (t *TV) NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error) {
	const op = "tv.NextChannel"
	log := t.logger.With("method", op)
	log.Info("Getting Channels")

	next := true
	back := false
	if boolStr == "false" {
		next = false
		back = true
	}
	channel, err := t.client.TvService().PreviousAndNext(ctx, &tv1.PreviousAndNextReq{
		UserID: id,
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

func (t *TV) AddChannel(ctx context.Context, req *entity.AddChannelReq) (*entity.StatusMessage, error) {
	const op = "tv.AddChannel"
	log := t.logger.With("method", op)
	log.Info("Adding Channel")

	addReq := entity.AddChannelReqPub{
		UserID:      req.UserID,
		ChannelName: req.ChannelName,
		MethodName:  "TV.AddChannel",
	}
	reqBody, err := json.Marshal(addReq)
	if err != nil {
		log.Error(err.Error())
	}
	err = t.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Added Channel successfully")
	return &entity.StatusMessage{
		Message: "Added Channel in processing",
	}, nil
}

func (t *TV) DeleteChannel(ctx context.Context, req *entity.DeleteChannelReq) (*entity.StatusMessage, error) {
	const op = "tv.DeleteChannel"
	log := t.logger.With("method", op)
	log.Info("Delete channel")
	deleteReq := entity.DeleteChannelReqPub{
		UserID:      req.UserID,
		ChannelName: req.ChannelName,
		MethodName:  "TV.DeleteChannel",
	}
	reqBody, err := json.Marshal(deleteReq)
	if err != nil {
		log.Error(err.Error())
	}
	err = t.p.Pub(ctx, reqBody, "devices")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Deleted Channel successfully")
	return &entity.StatusMessage{
		Message: "Deleted Channel in processing",
	}, nil
}

func (t *TV) UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceTvRes, error) {
	const op = "tv.UpVoice"
	log := t.logger.With("method", op)
	up := true
	down := false
	if boolStr == "false" {
		up = false
		down = true
	}
	res, err := t.client.TvService().DownOrUpVVoiceTv(ctx, &tv1.DownOrUpVVoiceTvReq{
		UserID: userID,
		Down:   down,
		Up:     up,
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Up Voice successfully")
	return &entity.DownOrUpVoiceTvRes{
		Sound: res.Sound,
	}, nil
}
