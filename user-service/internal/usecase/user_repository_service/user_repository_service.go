package user_repository_service

import (
	"context"
	"user_service_smart_home/internal/entity"
)

type (
	UserSaver interface {
		SaverUser(ctx context.Context, req *entity.CreateUserReq) error
	}
	ProviderUserRepositoryService interface {
		GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error)
		GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error)
	}
	UpdaterUserRepositoryService interface {
		UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error
		UpdateUserPassword(ctx context.Context, req *entity.UpdatePasswordReq) error
		UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error
	}
	DeleterUserRepositoryService interface {
		DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error
	}
)
type UserRepoUseCase struct {
	saver    UserSaver
	provider ProviderUserRepositoryService
	updater  UpdaterUserRepositoryService
	deleter  DeleterUserRepositoryService
}

func NewUserRepositoryService(saver UserSaver,
	provider ProviderUserRepositoryService,
	Updater UpdaterUserRepositoryService,
	Deleter DeleterUserRepositoryService) *UserRepoUseCase {
	return &UserRepoUseCase{
		saver:    saver,
		provider: provider,
		updater:  Updater,
		deleter:  Deleter,
	}
}
func (u *UserRepoUseCase) GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error) {
	return u.provider.GetUser(ctx, req)
}
func (u *UserRepoUseCase) GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error) {
	return u.provider.GetAllUser(ctx, req)
}
func (u *UserRepoUseCase) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error {
	return u.updater.UpdateUser(ctx, req)
}
func (u *UserRepoUseCase) UpdateUserPassword(ctx context.Context, req *entity.UpdatePasswordReq) error {
	return u.updater.UpdateUserPassword(ctx, req)
}
func (u *UserRepoUseCase) UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error {
	return u.updater.UpdateUserEmail(ctx, req)
}
func (u *UserRepoUseCase) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	return u.deleter.DeleteUser(ctx, req)
}
func (u *UserRepoUseCase) SaverUser(ctx context.Context, req *entity.CreateUserReq) error {
	return u.saver.SaverUser(ctx, req)
}
