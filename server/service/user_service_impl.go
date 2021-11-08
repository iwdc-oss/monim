package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/iwdc-oss/monim/server/helper"
	"github.com/iwdc-oss/monim/server/model/domain"
	"github.com/iwdc-oss/monim/server/model/web"
	"github.com/iwdc-oss/monim/server/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username:     request.Username,
		Email:        request.Email,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Password:     request.Password,
		ProfileImage: request.ProfileImage,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)

}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	return web.UserResponse{}
}

func (service *UserServiceImpl) FindById(ctx context.Context, userID int) web.UserResponse {
	return web.UserResponse{}

}

func (service *UserServiceImpl) FindByUsernameAndPassword(ctx context.Context, request web.UserFindByUsernameAndPasswordRequest) web.UserResponse {
	return web.UserResponse{}
}
