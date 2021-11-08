package service

import (
	"context"

	"github.com/iwdc-oss/monim/server/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	FindById(ctx context.Context, userID int) web.UserResponse
	FindByUsernameAndPassword(ctx context.Context, request web.UserFindByUsernameAndPasswordRequest) web.UserResponse
}
