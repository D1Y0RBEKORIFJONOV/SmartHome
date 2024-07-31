package user_tv_service_usecase

import (
	"context"
	tv_entity "device_service/internal/entity/tv"
)

type TVManagementUseCase interface {
	AddTvToUser(ctx context.Context, req *tv_entity.AddTVReq) error
	AddChannel(ctx context.Context, req *tv_entity.AddChannelReq) error
	GetUserChannel(ctx context.Context, req *tv_entity.GetUserChannelReq) (*tv_entity.GetUserChannelRes, error)
	DeleteChannel(ctx context.Context, req *tv_entity.DeleteChannelReq) error
	DownOrUpVVoiceTv(ctx context.Context, req *tv_entity.DownOrUpVoiceTvReq) (*tv_entity.DownOrUpVoiceTvRes, error)
	PreviousAndNext(ctx context.Context, req *tv_entity.PreviousAndNextReq) (*tv_entity.PreviousAndNextRes, error)
	OnAndOffUsersTv(ctx context.Context, req *tv_entity.OnAndOfReq) (string, error)
}

type tvManagementUseCase struct {
	TV TVManagementUseCase
}

func NewTVManagementUseCase(tv TVManagementUseCase) TVManagementUseCase {
	return &tvManagementUseCase{TV: tv}
}

func (u *tvManagementUseCase) AddTvToUser(ctx context.Context, req *tv_entity.AddTVReq) error {
	return u.TV.AddTvToUser(ctx, req)
}

func (u *tvManagementUseCase) AddChannel(ctx context.Context, req *tv_entity.AddChannelReq) error {
	return u.TV.AddChannel(ctx, req)
}

func (u *tvManagementUseCase) DeleteChannel(ctx context.Context, req *tv_entity.DeleteChannelReq) error {
	return u.TV.DeleteChannel(ctx, req)
}

func (u *tvManagementUseCase) GetUserChannel(ctx context.Context, req *tv_entity.GetUserChannelReq) (*tv_entity.GetUserChannelRes, error) {
	return u.TV.GetUserChannel(ctx, req)
}
func (u *tvManagementUseCase) DownOrUpVVoiceTv(ctx context.Context, req *tv_entity.DownOrUpVoiceTvReq) (*tv_entity.DownOrUpVoiceTvRes, error) {
	return u.TV.DownOrUpVVoiceTv(ctx, req)
}
func (u *tvManagementUseCase) PreviousAndNext(ctx context.Context, req *tv_entity.PreviousAndNextReq) (*tv_entity.PreviousAndNextRes, error) {
	return u.TV.PreviousAndNext(ctx, req)
}

func (u *tvManagementUseCase) OnAndOffUsersTv(ctx context.Context, req *tv_entity.OnAndOfReq) (string, error) {
	return u.TV.OnAndOffUsersTv(ctx, req)
}
