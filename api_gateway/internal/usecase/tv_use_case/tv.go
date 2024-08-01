package tv_use_case

import (
	"api_gate_way/internal/entity"
	"context"
)

//service TVService {
//rpc AddTV(AddTVReq) returns (TvStatusMessage);
//rpc AddChannel(AddChannelReq) returns (TvStatusMessage);
//rpc GetUserChannel(GetUserChannelReq) returns (GetUserChannelRes);
//rpc DeleteChannel(DeleteChannelReq) returns (TvStatusMessage);
//rpc DownOrUpVVoiceTv(DownOrUpVVoiceTvReq) returns (DownOrUpVVoiceTvRes);
//rpc PreviousAndNext(PreviousAndNextReq) returns (PreviousAndNextRes);
//rpc OnAndOffUserTv(OnAndOffUserTvReq) returns (OnAndOffUserTvRes);
//}

type TvUseCase interface {
	AddTV(ctx context.Context, req *entity.AddTVReq) (*entity.StatusMessage, error)
	OpenTV(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error)
	GetChannels(ctx context.Context, id string) (*entity.GetUserChannelRes, error)
	NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error)
	AddChannel(ctx context.Context, req *entity.AddChannelReq) (*entity.StatusMessage, error)
	DeleteChannel(ctx context.Context, req *entity.DeleteChannelReq) (*entity.StatusMessage, error)
	UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceTvRes, error)
}

type TvUseCaseImpl struct {
	tv TvUseCase
}

func NewTvUseCase(tv TvUseCase) *TvUseCaseImpl {
	return &TvUseCaseImpl{tv: tv}
}

func (t *TvUseCaseImpl) AddTV(ctx context.Context, req *entity.AddTVReq) (*entity.StatusMessage, error) {
	return t.tv.AddTV(ctx, req)
}

func (t *TvUseCaseImpl) OpenTV(ctx context.Context, boolStr string, id string) (*entity.StatusMessage, error) {
	return t.tv.OpenTV(ctx, boolStr, id)
}

func (t *TvUseCaseImpl) GetChannels(ctx context.Context, id string) (*entity.GetUserChannelRes, error) {
	return t.tv.GetChannels(ctx, id)
}

func (t *TvUseCaseImpl) NextChannel(ctx context.Context, id, boolStr string) (*entity.PreviousAndNextRes, error) {
	return t.tv.NextChannel(ctx, id, boolStr)
}

func (t *TvUseCaseImpl) AddChannel(ctx context.Context, req *entity.AddChannelReq) (*entity.StatusMessage, error) {
	return t.tv.AddChannel(ctx, req)
}

func (t *TvUseCaseImpl) DeleteChannel(ctx context.Context, req *entity.DeleteChannelReq) (*entity.StatusMessage, error) {
	return t.tv.DeleteChannel(ctx, req)
}

func (t *TvUseCaseImpl) UpVoice(ctx context.Context, userID, boolStr string) (*entity.DownOrUpVoiceTvRes, error) {
	return t.tv.UpVoice(ctx, userID, boolStr)
}
