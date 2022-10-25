package service

import (
	"context"
	"web_backend/internal/model"
	"web_backend/internal/model/entity"
	// "web_backend/internal/model/entity"
)

type IUser interface {
	Create(ctx context.Context, in model.UserCreateInput) (err error)
	IsSignedIn(ctx context.Context) bool
	SignIn(ctx context.Context, in model.UserSignInInput) (user *entity.User, err error)
	SignOut(ctx context.Context) error
	IsPassportAvailable(ctx context.Context, passport string) (bool, error)
	IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
	GetProfile(ctx context.Context) *entity.User
	GetUserList(ctx context.Context) (userList []*entity.User, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
