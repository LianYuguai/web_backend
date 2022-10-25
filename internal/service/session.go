// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"

	"web_backend/internal/model/entity"
)

type ISession interface {
	SetUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context) *entity.User
	RemoveUser(ctx context.Context) error
}

var localSession ISession

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}