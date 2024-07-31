package tv_server

import (
	"context"
	tv_entity "device_service/internal/entity/tv"
	"device_service/internal/usecase/tv_usecase/tv"
	"errors"
	"fmt"
	tv1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/TV"
	"google.golang.org/grpc"
)

type TVServer struct {
	tv1.UnimplementedTVServiceServer
	tv user_tv_service_usecase.TVManagementUseCase
}

func RegisterTVServer(server *grpc.Server, tv user_tv_service_usecase.TVManagementUseCase) {
	tv1.RegisterTVServiceServer(server, &TVServer{
		tv: tv,
	})
}

func (s *TVServer) AddTV(ctx context.Context, req *tv1.AddTVReq) (*tv1.TvStatusMessage, error) {
	if req.UserId == "" || req.ModelName == "" {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, errors.New("user id or model name is empty")
	}

	err := s.tv.AddTvToUser(ctx, &tv_entity.AddTVReq{
		UserID:        req.UserId,
		ModelName:     req.ModelName,
		CursorChannel: 1,
		Sound:         35,
	})
	if err != nil {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, err
	}

	return &tv1.TvStatusMessage{
		Successfully: true,
	}, nil
}

func (s *TVServer) AddChannel(ctx context.Context, req *tv1.AddChannelReq) (*tv1.TvStatusMessage, error) {
	if req.UserId == "" || req.ChannelName == "" {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, errors.New("user id or channel name is empty")
	}

	err := s.tv.AddChannel(ctx, &tv_entity.AddChannelReq{
		UserID:      req.UserId,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, err
	}
	return &tv1.TvStatusMessage{
		Successfully: true,
	}, nil
}

func (s *TVServer) GetUserChannel(ctx context.Context, req *tv1.GetUserChannelReq) (*tv1.GetUserChannelRes, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is empty")
	}
	resp, err := s.tv.GetUserChannel(ctx, &tv_entity.GetUserChannelReq{
		UserID:      req.UserId,
		CurrentName: req.ChannelName,
	})
	if err != nil {
		return &tv1.GetUserChannelRes{}, err
	}
	var res tv1.GetUserChannelRes
	for _, channel := range resp.Channels {
		res.Channels = append(res.Channels, &tv1.Channel{
			ChannelName:   channel.ChannelName,
			ChannelNumber: channel.ChannelNumber,
		})
		res.Count += 1
	}

	return &res, nil
}

func (s *TVServer) DeleteChannel(ctx context.Context, req *tv1.DeleteChannelReq) (*tv1.TvStatusMessage, error) {
	if req.UserId == "" || req.ChannelName == "" {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, errors.New("user id or channel name is empty")
	}
	err := s.tv.DeleteChannel(ctx, &tv_entity.DeleteChannelReq{
		UserID:      req.UserId,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		return &tv1.TvStatusMessage{
			Successfully: false,
		}, nil
	}
	return &tv1.TvStatusMessage{
		Successfully: true,
	}, nil
}

func (s *TVServer) DownOrUpVVoiceTv(ctx context.Context, req *tv1.DownOrUpVVoiceTvReq) (*tv1.DownOrUpVVoiceTvRes, error) {
	reap, err := s.tv.DownOrUpVVoiceTv(ctx, &tv_entity.DownOrUpVoiceTvReq{
		Down:   req.Down,
		Up:     req.Up,
		UserID: req.UserID,
	})
	if err != nil {
		return &tv1.DownOrUpVVoiceTvRes{}, err
	}

	return &tv1.DownOrUpVVoiceTvRes{
		Sound: reap.Sound,
	}, nil
}

func (s *TVServer) PreviousAndNext(ctx context.Context, req *tv1.PreviousAndNextReq) (*tv1.PreviousAndNextRes, error) {
	resp, err := s.tv.PreviousAndNext(ctx, &tv_entity.PreviousAndNextReq{
		Next:   req.Next,
		Back:   req.Back,
		UserID: req.UserID,
	})
	if err != nil {
		return &tv1.PreviousAndNextRes{}, err
	}

	if resp == nil || resp.Channel == nil {
		return nil, fmt.Errorf("received nil response or channel")
	}

	var channel tv1.Channel
	channel.ChannelNumber = resp.Channel.ChannelNumber
	channel.ChannelName = resp.Channel.ChannelName
	return &tv1.PreviousAndNextRes{
		Channel: &channel,
	}, nil
}

func (s *TVServer) OnAndOffUserTv(ctx context.Context, req *tv1.OnAndOffUserTvReq) (*tv1.OnAndOffUserTvRes, error) {
	message, err := s.tv.OnAndOffUsersTv(ctx, &tv_entity.OnAndOfReq{
		UserID: req.UserId,
		On:     req.On,
		Off:    req.Off,
	})
	if err != nil {
		return &tv1.OnAndOffUserTvRes{}, err
	}
	return &tv1.OnAndOffUserTvRes{
		Message: message,
	}, nil
}
