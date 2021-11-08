package helper

import (
	"github.com/iwdc-oss/monim/server/model/domain"
	"github.com/iwdc-oss/monim/server/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		ProfileImage: user.ProfileImage,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
