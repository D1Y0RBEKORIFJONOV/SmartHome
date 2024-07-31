package speaker_server

import (
	"context"
	speaker_entity "device_service/internal/entity/speaker"
	"device_service/internal/usecase/speaker_usecase/speaker"
	"errors"
	"fmt"
	speaker1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/speaker"

	"google.golang.org/grpc"
)

type SpeakerServer struct {
	speaker1.UnimplementedSpeakerServiceServer
	speaker user_speaker_service_usecase.SpeakerManagementUseCase
}

func RegisterSpeakerServer(server *grpc.Server, speaker user_speaker_service_usecase.SpeakerManagementUseCase) {
	speaker1.RegisterSpeakerServiceServer(server, &SpeakerServer{
		speaker: speaker,
	})
}

func (s *SpeakerServer) AddSpeaker(ctx context.Context, req *speaker1.AddSpeakerReq) (*speaker1.SpeakerStatusMessage, error) {
	if req.UserId == "" || req.ModelName == "" {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, errors.New("user id or model name is empty")
	}

	err := s.speaker.AddSpeakerToUser(ctx, &speaker_entity.AddSpeakerReq{
		UserID:        req.UserId,
		ModelName:     req.ModelName,
		CursorChannel: 1,
		Sound:         35,
	})
	if err != nil {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, err
	}

	return &speaker1.SpeakerStatusMessage{
		Successfully: true,
	}, nil
}

func (s *SpeakerServer) AddChannel(ctx context.Context, req *speaker1.AddChannelReqS) (*speaker1.SpeakerStatusMessage, error) {
	if req.UserId == "" || req.ChannelName == "" {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, errors.New("user id or channel name is empty")
	}

	err := s.speaker.AddChannel(ctx, &speaker_entity.AddChannelReq{
		UserID:      req.UserId,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, err
	}
	return &speaker1.SpeakerStatusMessage{
		Successfully: true,
	}, nil
}

func (s *SpeakerServer) GetUserChannel(ctx context.Context, req *speaker1.GetUserChannelReqS) (*speaker1.GetUserChannelResS, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is empty")
	}
	resp, err := s.speaker.GetUserChannel(ctx, &speaker_entity.GetUserChannelReq{
		UserID:      req.UserId,
		CurrentName: req.ChannelName,
	})
	if err != nil {
		return &speaker1.GetUserChannelResS{}, err
	}
	var res speaker1.GetUserChannelResS
	for _, channel := range resp.Channels {
		res.Channels = append(res.Channels, &speaker1.ChannelS{
			ChannelName:   channel.ChannelName,
			ChannelNumber: channel.ChannelNumber,
		})
		res.Count += 1
	}

	return &res, nil
}

func (s *SpeakerServer) DeleteChannel(ctx context.Context, req *speaker1.DeleteChannelReqS) (*speaker1.SpeakerStatusMessage, error) {
	if req.UserId == "" || req.ChannelName == "" {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, errors.New("user id or channel name is empty")
	}
	err := s.speaker.DeleteChannel(ctx, &speaker_entity.DeleteChannelReq{
		UserID:      req.UserId,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		return &speaker1.SpeakerStatusMessage{
			Successfully: false,
		}, nil
	}
	return &speaker1.SpeakerStatusMessage{
		Successfully: true,
	}, nil
}

func (s *SpeakerServer) DownOrUpVolume(ctx context.Context, req *speaker1.DownOrUpVolumeReqS) (*speaker1.DownOrUpVolumeResS, error) {
	reap, err := s.speaker.DownOrUpVolume(ctx, &speaker_entity.DownOrUpVolumeReq{
		Down:   req.Down,
		Up:     req.Up,
		UserID: req.UserId,
	})
	if err != nil {
		return &speaker1.DownOrUpVolumeResS{}, err
	}

	return &speaker1.DownOrUpVolumeResS{
		Sound: reap.Sound,
	}, nil
}

func (s *SpeakerServer) PreviousAndNext(ctx context.Context, req *speaker1.PreviousAndNextReqS) (*speaker1.PreviousAndNextResS, error) {
	resp, err := s.speaker.PreviousAndNext(ctx, &speaker_entity.PreviousAndNextReq{
		Next:   req.Next,
		Back:   req.Back,
		UserID: req.UserId,
	})
	if err != nil {
		return &speaker1.PreviousAndNextResS{}, err
	}

	if resp == nil || resp.Channel == nil {
		return nil, fmt.Errorf("received nil response or channel")
	}

	var channel speaker1.ChannelS
	channel.ChannelNumber = resp.Channel.ChannelNumber
	channel.ChannelName = resp.Channel.ChannelName
	return &speaker1.PreviousAndNextResS{
		Channel: &channel,
	}, nil
}

func (s *SpeakerServer) OnAndOffUserSpeaker(ctx context.Context, req *speaker1.OnAndOffUserSpeakerReq) (*speaker1.OnAndOffUserSpeakerRes, error) {
	message, err := s.speaker.OnAndOffUsersSpeaker(ctx, &speaker_entity.OnAndOffReq{
		UserID: req.UserId,
		On:     req.On,
		Off:    req.Off,
	})
	if err != nil {
		return &speaker1.OnAndOffUserSpeakerRes{}, err
	}
	return &speaker1.OnAndOffUserSpeakerRes{
		Message: message,
	}, nil
}
