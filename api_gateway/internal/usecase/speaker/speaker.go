package speaker_use_case

import (
	"api_gate_way/internal/entity"
	"context"
)

type SpeakerUseCase interface {
	AddSpeaker(ctx context.Context, req *entity.AddSpeakerReq) (*entity.StatusMessage, error)
	OpenSpeaker(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error)
	GetChannels(ctx context.Context, id string) (*entity.GetUserSongsRes, error)
	NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error)
	AddChannel(ctx context.Context, req *entity.AddChannelReqChannel) (*entity.StatusMessage, error)
	DeleteChannel(ctx context.Context, req *entity.DeleteChannelReqChannel) (*entity.StatusMessage, error)
	UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceSpeakerRes, error)
}

type SpeakerUseCaseImpl struct {
	speaker SpeakerUseCase
}

func NewSpeakerUseCase(speaker SpeakerUseCase) *SpeakerUseCaseImpl {
	return &SpeakerUseCaseImpl{speaker: speaker}
}

func (t *SpeakerUseCaseImpl) AddSpeaker(ctx context.Context, req *entity.AddSpeakerReq) (*entity.StatusMessage, error) {
	return t.speaker.AddSpeaker(ctx, req)
}
func (t *SpeakerUseCaseImpl) OpenSpeaker(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error) {
	return t.speaker.OpenSpeaker(ctx, boolStr, id)
}
func (t *SpeakerUseCaseImpl) GetChannels(ctx context.Context, id string) (*entity.GetUserSongsRes, error) {
	return t.speaker.GetChannels(ctx, id)

}
func (t *SpeakerUseCaseImpl) AddChannel(ctx context.Context, req *entity.AddChannelReqChannel) (*entity.StatusMessage, error) {
	return t.speaker.AddChannel(ctx, req)

}
func (t *SpeakerUseCaseImpl) DeleteChannel(ctx context.Context, req *entity.DeleteChannelReqChannel) (*entity.StatusMessage, error) {
	return t.speaker.DeleteChannel(ctx, req)

}
func (t *SpeakerUseCaseImpl) UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceSpeakerRes, error) {
	return t.speaker.UpVoice(ctx, userID, boolStr)
}
func (t *SpeakerUseCaseImpl) NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error) {
	return t.speaker.NextChannel(ctx, id, boolStr)
}
