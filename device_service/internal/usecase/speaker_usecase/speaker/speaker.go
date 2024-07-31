package user_speaker_service_usecase

import (
	"context"
	speaker_entity "device_service/internal/entity/speaker"
)

type SpeakerManagementUseCase interface {
	AddSpeakerToUser(ctx context.Context, req *speaker_entity.AddSpeakerReq) error
	AddChannel(ctx context.Context, req *speaker_entity.AddChannelReq) error
	GetUserChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error)
	DeleteChannel(ctx context.Context, req *speaker_entity.DeleteChannelReq) error
	DownOrUpVolume(ctx context.Context, req *speaker_entity.DownOrUpVolumeReq) (*speaker_entity.DownOrUpVolumeRes, error)
	PreviousAndNext(ctx context.Context, req *speaker_entity.PreviousAndNextReq) (*speaker_entity.PreviousAndNextRes, error)
	OnAndOffUsersSpeaker(ctx context.Context, req *speaker_entity.OnAndOffReq) (string, error)
}

type speakerManagementUseCase struct {
	Speaker SpeakerManagementUseCase
}

func NewSpeakerManagementUseCase(speaker SpeakerManagementUseCase) SpeakerManagementUseCase {
	return &speakerManagementUseCase{Speaker: speaker}
}

func (u *speakerManagementUseCase) AddSpeakerToUser(ctx context.Context, req *speaker_entity.AddSpeakerReq) error {
	return u.Speaker.AddSpeakerToUser(ctx, req)
}

func (u *speakerManagementUseCase) AddChannel(ctx context.Context, req *speaker_entity.AddChannelReq) error {
	return u.Speaker.AddChannel(ctx, req)
}

func (u *speakerManagementUseCase) DeleteChannel(ctx context.Context, req *speaker_entity.DeleteChannelReq) error {
	return u.Speaker.DeleteChannel(ctx, req)
}

func (u *speakerManagementUseCase) GetUserChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error) {
	return u.Speaker.GetUserChannel(ctx, req)
}

func (u *speakerManagementUseCase) DownOrUpVolume(ctx context.Context, req *speaker_entity.DownOrUpVolumeReq) (*speaker_entity.DownOrUpVolumeRes, error) {
	return u.Speaker.DownOrUpVolume(ctx, req)
}

func (u *speakerManagementUseCase) PreviousAndNext(ctx context.Context, req *speaker_entity.PreviousAndNextReq) (*speaker_entity.PreviousAndNextRes, error) {
	return u.Speaker.PreviousAndNext(ctx, req)
}

func (u *speakerManagementUseCase) OnAndOffUsersSpeaker(ctx context.Context, req *speaker_entity.OnAndOffReq) (string, error) {
	return u.Speaker.OnAndOffUsersSpeaker(ctx, req)
}
