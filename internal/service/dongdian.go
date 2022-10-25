package service

import (
	"context"
	"web_backend/internal/model/entity"
)

type IDongdian interface {
	GetG2g(ctx context.Context) (list []*entity.G2G, err error)
	GetGoldenchamData(ctx context.Context) (list []*entity.GoldenchamData, err error)
	GetGoldencham(ctx context.Context) (list []*entity.Goldencham, err error)
	GetOer(ctx context.Context) (list []*entity.Oer, err error)
	GetOoc(ctx context.Context) (list []*entity.Ooc, err error)
	GetPm(ctx context.Context) (list []*entity.Pm, err error)
	GetSeasoning(ctx context.Context) (list []*entity.Seasoning, err error)
}

var localDongdian IDongdian

func Dongdian() IDongdian {
	if localDongdian == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localDongdian
}

func RegisterDongdian(i IDongdian) {
	localDongdian = i
}
