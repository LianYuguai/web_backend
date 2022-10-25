package service

import (
	"context"
	"web_backend/internal/model/entity"
	// "web_backend/internal/model/entity"
)

type ILog interface {
	GetLogList(ctx context.Context) (logList []*entity.Log, err error)
}

var localLog ILog

func Log() ILog {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localLog
}

func RegisterLog(i ILog) {
	localLog = i
}
